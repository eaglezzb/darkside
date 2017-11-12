package model

import (
	"github.com/flywithbug/darkside/db"
	log "github.com/flywithbug/log4go"
	"fmt"
)

type SMSTXModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	SMStype 	string 		`json:"type,omitempty" form:"type,omitempty"`
	Messag		string 		`json:"msg,omitempty" form:"message,omitempty"`
	Result		int  		`json:"result,omitempty" form:"result,omitempty"`
	Time		int64  		`json:"time,omitempty" form:"time,omitempty"`
	Ext		string		`json:"ext,omitempty" form:"extra,omitempty"`
	Mobile		string 		`json:"mobile,omitempty" form:"mobile,omitempty"`
	Ncode		string 		`json:"nationcode,omitempty" form:"ncode,omitempty"`
	Errmsg		string		`json:"errmsg,omitempty" form:"errmsg,omitempty"`
	Sid		string		`json:"sid,omitempty" form:"sid,omitempty"`
	Fee		int		`json:"fee,omitempty" form:"fee,omitempty"`
	Smscode		string		`json:"smscode,omitempty" form:"smscode,omitempty"`
	TelModel	TelephoneModel	`json:"tel,omitempty" form:"tel,omitempty"`
}

type TelephoneModel struct {
	Code		string 			`json:"code,omitempty" form:"ncode,omitempty"`
	Mobile		string 			`json:"mobile,omitempty" form:"mobile,omitempty"`
}


func (sms *SMSTXModel)InsertSMSInfo()error {
	fmt.Println(sms)
	db := db.DBConf()
	stmt, err := db.Prepare("INSERT smstx SET type=?,message=?,result=?,time=?,ext=?,mobile=?,ncode=?,errmsg=?,sid=?,fee=?,smscode=?")
	checkSMSErr(err)
	_, err = stmt.Exec(sms.SMStype,sms.Messag,sms.Result,sms.Time,sms.Ext,sms.Mobile,sms.Ncode,sms.Errmsg,sms.Sid,sms.Fee,sms.Smscode)
	checkSMSErr(err)

	//手机数据短信send次数自增
	stmt, err = db.Prepare("INSERT telephone SET mobile=?,ncode=? on duplicate key update scount=scount+1,ncode=?")
	checkSMSErr(err)
	_, err1 := stmt.Exec(sms.TelModel.Mobile,sms.TelModel.Code,sms.TelModel.Code)
	checkSMSErr(err1)
	return err
}

//func FindSMSRecBy()  {
//
//}




func checkSMSErr(err error) {
	if err != nil {
		log.Warn(err.Error())
	}
}