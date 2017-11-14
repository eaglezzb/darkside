package model

import (
	"github.com/flywithbug/darkside/db"
	log "github.com/flywithbug/log4go"
	"errors"
	"fmt"
)

type BaseUserInfoModel struct {

}

type UserInfoModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	UserName	string 		`json:"username,omitempty" form:"username,omitempty"`
	Password	string  	`json:"password,omitempty" form:"password,omitempty"`
	CreateTime	int64  		`json:"createtime,omitempty" form:"createtime,omitempty"`
	UpdateTime	int64  		`json:"updatetime,omitempty" form:"updatetime,omitempty"`
	Sex 		int 		`json:"sex,omitempty" form:"sex,omitempty"`     //0默认未设置 1男，2女
	UserId 		string 		`json:"userid,omitempty" form:"userid,omitempty"`
	DepartName	string		`json:"departname,omitempty" form:"departname,omitempty"`
	Phone 		string		`json:"phone,omitempty" form:"phone,omitempty"`
	PhonePrefix 	string		`json:"phoneprefix,omitempty" form:"phoneprefix,omitempty"`
	Mail 		string		`json:"mail,omitempty" form:"mail,omitempty"`
	OldPassword 	string		`json:"oldpassword,omitempty" form:"oldpassword,omitempty"`
	Authtoken 	string		`json:"authtoken,omitempty" form:"authtoken,omitempty"`
	State 		int 		`json:"state,omitempty" form:"state,omitempty"`

	RegisterType 	string          `json:"registertype,omitempty"`
	VerifyCode      string 		`json:"verifycode,omitempty"`
}




func NewUser() UserInfoModel {
	return UserInfoModel{}
}

func (user *UserInfoModel)ToString()(desc string)  {
	desc = "name:"+user.UserName
	return desc
}

func (user *UserInfoModel)InsertUser()(error){
	if len(user.UserName) >0 && CheckUserNameValid(user.UserName) == false{
		err := errors.New("username already exists")
		return err
	}

	if CheckPhoneValid(user.Phone) == false{
		err := errors.New("phone number already exists")
		return err
	}

	db := db.DBConf()
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?,updatetime=?,password=?,sex=?,mail=?,phone=?,phoneprefix=?,state=?")
	checkErr(err)
	_, err = stmt.Exec(user.UserName, user.DepartName,user.CreateTime,user.UpdateTime,user.Password,user.Sex,user.Mail,user.Phone,user.PhonePrefix,user.State)
	checkErr(err)
	return err
}

func (user *UserInfoModel)UpdateIntoDB()(error)  {
	if user.Uid == 0 {
		log.Error("update faild，Pri key Uid not found")
		return errors.New("update faild，Pri key Uid not found")
	}
	db := db.DBConf()
	stmt, err := db.Prepare("UPDATE userinfo set username=?,departname=?,createtime=?,password=?,sex=? where uid=?")
	checkErr(err)
	_, err = stmt.Exec(user.UserName, user.DepartName,user.CreateTime,user.Password,user.Sex, user.Uid)
	checkErr(err)
	return err
}


func CheckUserNameValid(name string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE username=?", name).Scan(&name)
	if err != nil {
		return true
	}
	log.Warn("username already exist",name)
	return false
}

func CheckPhoneValid(phone string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  phone FROM userinfo WHERE phone=?", phone).Scan(&phone)
	if err != nil {
		return true
	}
	log.Warn("phone number already exists",phone)
	return false
}

func CheckEmailValid(mail string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE mail=?", mail).Scan(&mail)
	if err != nil {
		return true
	}
	log.Warn("mail already exists",mail)
	return false
}


func CheckUserIdValid(userId string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE userid=?", userId).Scan(&userId)
	if err != nil {
		return true
	}
	log.Warn("CheckUserIdValid：userid already exists",userId)
	return false
}

func FindUserFromDB(uid int64)(UserInfoModel,error)  {
	var user UserInfoModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, departname, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM userinfo WHERE uid=?", uid).
		Scan(&user.Uid,
		 &user.UserName, &user.DepartName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)
	checkErr(err)
	return user,err
}

func FindUserFromDBByName(name string)(UserInfoModel,error)  {
	var user UserInfoModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, departname, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM userinfo WHERE username=?", name).
		Scan(&user.Uid,
		&user.UserName, &user.DepartName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)
	user.Password = ""
	user.OldPassword = ""
	checkErr(err)
	return user,err
}

func CheckUserNameAndPass(name string,pass string)(UserInfoModel,error)  {
	var user UserInfoModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, departname, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM userinfo WHERE username=?", name).
		Scan(&user.Uid,
		&user.UserName, &user.DepartName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)

	if user.Password != pass {
		err = errors.New("password not right")
		user = NewUser()
	}
	user.Password = ""
	user.OldPassword = ""
	checkErr(err)
	return user,err
}


func DeleteUserFromDB(uid int64)(error)  {
	db := db.DBConf()
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	if err != nil{
		return err
	}
	fmt.Println(uid)
	_,err = stmt.Exec(uid)
	return err
}

func checkErr(err error) {
	if err != nil {
		log.Warn(err.Error())
	}
}