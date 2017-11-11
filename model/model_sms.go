package model

import (
	"github.com/flywithbug/darkside/db"

)

type SMSTXModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	SMStype 	int 		`json:"type,omitempty" form:"type,omitempty"`     //0默认未设置 1男，2女
	Messag		string 		`json:"msg,omitempty" form:"message,omitempty"`
	Signature	string  	`json:"sig,omitempty" form:"signature,omitempty"`
	Time		int64  		`json:"time,omitempty" form:"time,omitempty"`
	Extend 		string 		`json:"extend,omitempty" form:"extend,omitempty"`
	Ext		string		`json:"ext,omitempty" form:"extra,omitempty"`
	Mobile		int 		`json:"mobile,omitempty" form:"mobile,omitempty"`
	Ncode		int 		`json:"nationcode,omitempty" form:"ncode,omitempty"`
	TelModel	telephoneModel	`json:"tel,omitempty" form:"tel,omitempty"`
}

type telephoneModel struct {
	Code		int 			`json:"nationcode,omitempty" form:"ncode,omitempty"`
	Mobile		int 			`json:"mobile,omitempty" form:"mobile,omitempty"`
}

func (sms *SMSTXModel)InsertSMSInfo()error {

	db := db.DBConf()
	stmt, err := db.Prepare("INSERT smstx SET type=?,message=?,signature=?,time=?,extend=?,extra=?,mobile=?,ncode=?")
	checkErr(err)
	_, err = stmt.Exec(sms.SMStype,sms.Messag,sms.Signature,sms.Time,sms.Extend,sms.Ext,sms.Mobile,sms.Ncode)
	checkErr(err)

	stmt, err = db.Prepare("INSERT telephone SET mobile=?,ncode=?")
	checkErr(err)
	_, err1 := stmt.Exec(sms.TelModel.Mobile,sms.TelModel.Code)
	checkErr(err1)
	return err
}


