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
	Phone 		int64		`json:"phone" form:"phone"`
	PhonePrefix 	int64		`json:"phoneprefix" form:"phoneprefix"`
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
		err := errors.New("用户名已存在")
		return err
	}
	//if user.Mail != nil && CheckEmailValid(user.Mail) == false {
	//	err := errors.New("邮箱已存在")
	//	return err
	//}

	db := db.DBConf()
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?,updatetime=?,password=?,sex=?,mail=?")
	checkErr(err)
	_, err = stmt.Exec(user.UserName, user.DepartName,user.CreateTime,user.UpdateTime,user.Password,user.Sex,user.Mail)
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
	var username string
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE username=?", name).Scan(&username)
	log.Info("用户名已被注册",err)
	if len(username) == 0 {
		return true
	}
	return false
}

func CheckEmailValid(mail string)(bool)  {
	var username string
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE mail=?", mail).Scan(&username)
	log.Info("用户名可用",err)
	if len(username) == 0 {
		return true
	}
	return false
}


func CheckUserIdValid(userId string)(bool)  {
	var username string
	db := db.DBConf()
	err := db.QueryRow("SELECT  username FROM userinfo WHERE userid=?", userId).Scan(&username)
	log.Info("用户名可用",err)
	if len(username) == 0 {
		return true
	}
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
		log.Error(err.Error())
	}
}