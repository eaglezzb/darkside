package model

import (
	"github.com/brasbug/darkside/db"
	log "github.com/brasbug/log4go"
	"errors"
	"fmt"
)

type BaseUserInfoModel struct {

}

type UserInfoModel struct {
	Uid  		int64 		`json:"uid" form:"uid"`
	UserName	string 		`json:"username" form:"username"`
	Password	string  	`json:"password" form:"password"`
	CreateTime	int64  		`json:"createtime" form:"createtime"`
	UpdateTime	int64  		`json:"updatetime" form:"updatetime"`
	Sex 		int 		`json:"sex" form:"sex"`     //0默认未设置 1男，2女
	UserId 		string 		`json:"userid" form:"userid"`
	DepartName	string		`json:"departname" form:"departname"`
	Phone 		string		`json:"phone" form:"phone"`
	PhonePrefix 	string		`json:"phoneprefix" form:"phoneprefix"`
	Mail 		string		`json:"mail" form:"mail"`
	OldPassword 	string		`json:"oldpassword" form:"oldpassword"`
	Authtoken 	string		`json:"authtoken" form:"authtoken"`
	State 		int 		`json:"state" form:"state"`
}

func NewUser() UserInfoModel {
	return UserInfoModel{}
}

func (user *UserInfoModel)ToString()(desc string)  {
	desc = "name:"+user.UserName
	return desc
}

func (user *UserInfoModel)InsertUser()(error){

	if CheckUserNameValid(user.UserName) == false{
		err := errors.New("该用户名已被注册")
		return err
	}

	if CheckPhoneValid(user.Phone) == false{
		err := errors.New("该手机号已被注册")
		return err
	}

	db := db.DBConf()
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?,updatetime=?,password=?,sex=?,mail=?,phone=?,phoneprefix=?")
	checkErr(err)
	_, err = stmt.Exec(user.UserName, user.DepartName,user.CreateTime,user.UpdateTime,user.Password,user.Sex,user.Mail,user.Phone,user.PhonePrefix)
	checkErr(err)
	return err
}

func (user *UserInfoModel)UpdateIntoDB()(error)  {
	if user.Uid == 0 {
		log.Error("更新失败，找不到主键Uid")
		return errors.New("更新失败，找不到主键Uid")
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
	log.Warn("CheckUserNameValid：该用户名已存在",name)
	return false
}

func CheckPhoneValid(phone string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE phone=?", phone).Scan(&phone)
	if err != nil {
		return true
	}
	log.Warn("CheckPhoneValid：该手机号已被注册",phone)
	return false
}

func CheckEmailValid(mail string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE mail=?", mail).Scan(&mail)
	if err != nil {
		return true
	}
	log.Warn("CheckEmailValid：该邮箱已存在",mail)
	return false
}


func CheckUserIdValid(userId string)(bool)  {
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE userid=?", userId).Scan(&userId)
	if err != nil {
		return true
	}
	log.Warn("CheckUserIdValid：该用户Id已存在",userId)
	return false
}

func FindUserFromDB(uid int64)(UserInfoModel,error)  {
	var user UserInfoModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, departname, createtime, updatetime, sex, userId, phone, phoneprefix FROM userinfo WHERE uid=?", uid).Scan(&user.Uid,
		 &user.UserName, &user.DepartName, &user.CreateTime, &user.UpdateTime, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix)
	checkErr(err)
	return user,err
}

func FindUserFromDBByName(name string)(UserInfoModel,error)  {
	var user UserInfoModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, password, departname, createtime, updatetime, sex, userId, phone, phoneprefix FROM userinfo WHERE username=?", name).Scan(&user.Uid,
		&user.UserName, &user.Password, &user.DepartName, &user.CreateTime, &user.UpdateTime, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix)
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