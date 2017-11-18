package model

import (
	"github.com/flywithbug/darkside/db"
	log "github.com/flywithbug/log4go"
)

type TelephoneModel struct {
	NCode		string 			`json:"code,omitempty" form:"ncode,omitempty"`
	Mobile		string 			`json:"mobile,omitempty" form:"mobile,omitempty"`
	Type		int 			`json:"type,omitempty" form:"type,omitempty"` //1 用户注册类型
}

func (tel TelephoneModel)UpdateSendCount2DB()error {
	db := db.DBConf()
	//手机数据短信send次数自增
	stmt, err := db.Prepare("INSERT telephone SET mobile=?,ncode=? on duplicate key update scount=scount+1,ncode=?")
	if err != nil{
		log.Warn(err.Error())
		return err
	}
	_, err = stmt.Exec(tel.Mobile,tel.NCode,tel.NCode)
	if err != nil {
		log.Warn(err.Error())
	}
	return err
}