package model

import (
	"darkside/db"
	log "github.com/flywithbug/log4go"
	"fmt"
)

type EmailInfoModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	Mail 		string		`json:"mail,omitempty" form:"mail,omitempty"`
	Sender 		string		`json:"sender,omitempty" form:"sender,omitempty"`
	Type		int  		`json:"type,omitempty" form:"type,omitempty"`
	Verifycode	string 		`json:"verifycode,omitempty" form:"verifycode,omitempty"`
	Message		string  	`json:"message,omitempty" form:"message,omitempty"`
	CreateTime	int64  		`json:"createtime,omitempty" form:"createtime,omitempty"`
	Status          int 		`json:"status,omitempty" form:"status,omitempty"`
}



type MailModel struct {
	Email 		string		`json:"email,omitempty"`
	Type		int  		`json:"type,omitempty"`
}


func (email *EmailInfoModel)InsertSMSInfo()error {
	fmt.Println(email)
	db := db.DBConf()
	stmt, err := db.Prepare("INSERT mails SET mail=?,sender=?,verifycode=?,message=?,type=?,status=?,createtime=?")
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	_, err = stmt.Exec(email.Mail,email.Sender,email.Verifycode,email.Message,email.Type,email.Status,email.CreateTime)
	if err != nil {
		log.Warn(err.Error())
	}
	return err
}