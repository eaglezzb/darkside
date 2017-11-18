package model

import (
	"github.com/flywithbug/darkside/db"
	log "github.com/flywithbug/log4go"
	"errors"
	"fmt"
	"github.com/flywithbug/utils"
)

type BaseUserModel struct {

}

type UserModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	UserName	string 		`json:"username,omitempty" form:"username,omitempty"`
	Password	string  	`json:"password,omitempty" form:"password,omitempty"`
	CreateTime	int64  		`json:"createtime,omitempty" form:"createtime,omitempty"`
	UpdateTime	int64  		`json:"updatetime,omitempty" form:"updatetime,omitempty"`
	Sex 		int 		`json:"sex,omitempty" form:"sex,omitempty"`     //0默认未设置 1男，2女
	UserId 		string 		`json:"userid,omitempty" form:"userid,omitempty"`
	Phone 		string		`json:"phone,omitempty" form:"phone,omitempty"`
	PhonePrefix 	string		`json:"phoneprefix,omitempty" form:"phoneprefix,omitempty"`
	Mail 		string		`json:"mail,omitempty" form:"mail,omitempty"`
	OldPassword 	string		`json:"oldpassword,omitempty" form:"oldpassword,omitempty"`
	Authtoken 	string		`json:"authtoken,omitempty" form:"authtoken,omitempty"`
	State 		int 		`json:"state,omitempty" form:"state,omitempty"`

	RegisterType 	string          `json:"registertype,omitempty"`
	VerifyCode      string 		`json:"verifycode,omitempty"`
}




func NewUser() UserModel {
	return UserModel{}
}

func (user *UserModel)ToString()(desc string)  {
	desc = "name:"+user.UserName
	return desc
}

func (user *UserModel)InsertUser()(error){
	//if len(user.UserName) >0 && CheckUserNameValid(user.UserName) == false{
	//	err := errors.New("username already exists")
	//	return err
	//}

	if CheckPhoneValid(user.Phone) == false{
		err := errors.New("phone number already exists")
		return err
	}

	db := db.DBConf()
	stmt, err := db.Prepare("INSERT user SET username=?,createtime=?,updatetime=?,password=?,sex=?,mail=?,phone=?,phoneprefix=?,state=?,userid=?")
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	user.UserId = utils.Md5(user.Phone+user.UserName)
	_, err = stmt.Exec(user.UserName,user.CreateTime,user.UpdateTime,user.Password,user.Sex,user.Mail,user.Phone,user.PhonePrefix,user.State,user.UserId)
	if err != nil {
		log.Warn(err.Error())
	}
	return err
}

func (user *UserModel)UpdateIntoDB()(error)  {
	if user.Uid == 0 {
		log.Error("update faild，Pri key Uid not found")
		return errors.New("update faild，Pri key Uid not found")
	}
	db := db.DBConf()
	stmt, err := db.Prepare("UPDATE user set username=?,createtime=?,password=?,sex=? where uid=?")
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	_, err = stmt.Exec(user.UserName,user.CreateTime,user.Password,user.Sex, user.Uid)
	if err != nil {
		log.Warn(err.Error())
	}
	return err
}


func CheckUserNameValid(name string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM user WHERE username=?", name).Scan(&name)
	if err != nil {
		return true
	}
	log.Warn("username already exist",name)
	return false
}

func CheckPhoneValid(phone string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  phone FROM user WHERE phone=?", phone).Scan(&phone)
	if err != nil {
		return true
	}
	log.Warn("phone number already exists",phone)
	return false
}

func CheckEmailValid(mail string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM user WHERE mail=?", mail).Scan(&mail)
	if err != nil {
		return true
	}
	log.Warn("mail already exists",mail)
	return false
}


func CheckUserIdValid(userId string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM user WHERE userid=?", userId).Scan(&userId)
	if err != nil {
		return true
	}
	log.Warn("CheckUserIdValid：userid already exists",userId)
	return false
}

func FindUserFromDB(uid int64)(UserModel,error)  {
	var user UserModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM user WHERE uid=?", uid).
		Scan(&user.Uid,
		 &user.UserName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)

	if err != nil {
		log.Warn(err.Error())
		err = errors.New("user not found")
	}
	user.Password = ""
	user.OldPassword = ""
	return user,err
}

func FindUserFromDBByUserid(userid string)(UserModel,error)  {
	var user UserModel
	db := db.DBConf()
	err := db.QueryRow("SELECT  username, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM user WHERE userid=?", userid).
		Scan(&user.UserName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)
	if err != nil {
		log.Warn(err.Error())
	}
	user.Password = ""
	user.OldPassword = ""
	return user,err
}


func FindUserFromDBByName(name string)(UserModel,error)  {
	var user UserModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM user WHERE username=?", name).
		Scan(&user.Uid,
		&user.UserName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)

	if err != nil {
		log.Warn(err.Error())
		err = errors.New("user not found")

	}
	user.Password = ""
	user.OldPassword = ""
	return user,err
}

func UserLogin(name string,pass string)(UserModel,error)  {
	var user UserModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail FROM user WHERE username=? and password=?", name,pass).
		Scan(&user.Uid,
		&user.UserName, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail)
	if err != nil{
		log.Warn(err.Error())
		err = errors.New("user not found")
	}
	user.Password = ""
	user.OldPassword = ""
	return user,err
}

func DeleteUserFromDB(uid int64)(error)  {
	db := db.DBConf()
	stmt, err := db.Prepare("delete from user where uid=?")
	if err != nil{
		log.Warn(err.Error())
		return err
	}
	fmt.Println(uid)
	_,err = stmt.Exec(uid)
	return err
}
