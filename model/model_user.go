package model

import (
	"time"
	"github.com/brasbug/darkside/db"
	log "github.com/brasbug/log4go"
	"fmt"
)

type UserInfoModel struct {
	Id  		int64 		`json:"id" form:"id"`
	Uid  		string 		`json:"uid" form:"uid"`
	UserName	string 		`json:"username" form:"username"`
	Password	string  	`json:"password" form:"password"`
	CreateTime	time.Time  	`json:"createtime" form:"createtime"`
	UpdateTime	time.Time  	`json:"updatetime" form:"updatetime"`
	Sex 		int 		`json:"sex" form:"sex"`     //0默认未设置 1男，2女
	UserId 		string 		`json:"userid" form:"userid"`
	DepartName	string		`json:"departname" form:"departname"`
}

func NewUser() UserInfoModel {
	return UserInfoModel{}
}

func (user *UserInfoModel)ToString()(desc string)  {
	desc = "name:"+user.UserName + " " + "Uid:"+user.Uid
	return desc
}

func (user *UserInfoModel)InsertUser(){
	db := db.DBConf()
	fmt.Println(user.UserName)

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?,password=?,sex=?")
	checkErr(err)

	_, err = stmt.Exec(user.UserName, user.DepartName, time.Now(),user.Password,user.Sex)
	checkErr(err)


}

func (user *UserInfoModel)UpdateIntoDB()  {
	db := db.DBConf()

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?,password=?")
	checkErr(err)

	res, err := stmt.Exec("码农", "技术部", "2017-11-06")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}




func checkErr(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}






