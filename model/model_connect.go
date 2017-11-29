package model

import (
	"darkside/db"
	log "github.com/flywithbug/log4go"

)

type UserConnectModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	Userid_1 	float32		`json:"userid_1,omitempty" form:"userid_1,omitempty"`
	Userid_2	float32  	`json:"userid_2,omitempty" form:"userid_2,omitempty"`
}

func (ucnn *UserConnectModel)AddConnect()error  {
	db := db.DBConf()
	stmt, err := db.Prepare("INSERT user_friend_connect SET userid_1=?,userid_2=?")
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	if ucnn.Userid_1 > ucnn.Userid_2 {
		ucnn.Userid_1,ucnn.Userid_2 = ucnn.Userid_2,ucnn.Userid_1
	}
	_, err = stmt.Exec(ucnn.Userid_1,ucnn.Userid_2)
	if err != nil {
		log.Warn(err.Error())
	}
	return err
}





