package model

import (
	"github.com/flywithbug/darkside/db"
	log "github.com/flywithbug/log4go"
	"fmt"
	"time"
)

type LocationModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	Latitude 	float32		`json:"latitude,omitempty" form:"latitude,omitempty"`
	Longitude	float32  		`json:"longitude,omitempty" form:"longitude,omitempty"`
	Updatetime	int64  		`json:"time,omitempty" form:"time,omitempty"`
}

func (location *LocationModel)UpdateLocation()error {
	fmt.Println()
	db := db.DBConf()
	//手机数据短信send次数自增
	stmt, err := db.Prepare("INSERT location SET uid =?,latitude=?,longitude=?, time=? on duplicate key update uid=?, latitude=?,longitude=?, time=?")
	if err != nil{
		log.Warn(err.Error())
		return err
	}
	location.Updatetime = time.Now().Unix()
	_, err = stmt.Exec(location.Uid,location.Latitude,location.Longitude,location.Updatetime,location.Uid,location.Latitude,location.Longitude,location.Updatetime)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	stmt, err = db.Prepare("INSERT location_history_record SET uid =?,latitude=?,longitude=?, time=? ")
	_, err = stmt.Exec(location.Uid,location.Latitude,location.Longitude,location.Updatetime)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	return err
}