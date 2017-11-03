package model

import (
	"time"
	"github.com/brasbug/darkside/db"
	log "github.com/brasbug/log4go"
	"fmt"
)

type UserInfoModel struct {
	Id  		int64 	`json:"id" form:"id"`
	Uid  		string 	`json:"uid" form:"uid"`
	Name	 	string 	`json:"name" form:"name"`
	Password	string  `json:"password" form:"password"`
	CreateTime	time.Time  `json:"created" form:"createTime"`
	Sex 		string 	`json:"sex" form:"sex"`
	UserId 		string 	`json:"userid" form:"userid"`
}

func NewUser() UserInfoModel {
	return UserInfoModel{}
}

func (user *UserInfoModel)ToString()(desc string)  {
	desc = "name:"+user.Name + " " + "Uid:"+user.Uid
	return desc
}

func (user *UserInfoModel)InsertUser(){
	db := db.DBConf()

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?")
	checkErr(err)

	res, err := stmt.Exec("码农", "技术部", "2017-11-06")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

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






