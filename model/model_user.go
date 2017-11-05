package model

import (
	"time"
	"github.com/brasbug/darkside/db"
	log "github.com/brasbug/log4go"
	"fmt"
	"errors"
)

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
}


func NewUser() UserInfoModel {
	return UserInfoModel{}
}

func (user *UserInfoModel)ToString()(desc string)  {
	desc = "name:"+user.UserName
	return desc
}

func (user *UserInfoModel)InsertUser()(error){
	db := db.DBConf()
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?,password=?,sex=?")
	checkErr(err)
	_, err = stmt.Exec(user.UserName, user.DepartName, time.Now().Unix(),user.Password,user.Sex)
	checkErr(err)
	return err
}

func (user *UserInfoModel)UpdateIntoDB()(error)  {
	if user.Uid == 0 {
		log.Error("更新失败，找不到主键Uid")
		return errors.New("更新失败，找不到主键Uid")
	}
	db := db.DBConf()
	fmt.Println(user.Uid)
	stmt, err := db.Prepare("UPDATE userinfo set username=?,departname=?,createtime=?,password=?,sex=? where uid=?")
	checkErr(err)
	_, err = stmt.Exec(user.UserName, user.DepartName,user.CreateTime,user.Password,user.Sex, user.Uid)
	checkErr(err)
	return err
}

func FindUserFromDB(uid int64)(UserInfoModel,error)  {
	var user UserInfoModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, username, departname, createtime, updatetime, sex, userId, phone, phoneprefix FROM userinfo WHERE uid=?", uid).Scan(&user.Uid,
		 &user.UserName, &user.DepartName, &user.CreateTime, &user.UpdateTime, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix)
	fmt.Println("user.UserName",user.UserName)
	checkErr(err)
	return user,err
}

func DeleteUserFromDB(uid int64)(error)  {
	db := db.DBConf()
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	if err != nil{
		return err
	}
	_,err = stmt.Exec(uid)
	return err
}

func checkErr(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}