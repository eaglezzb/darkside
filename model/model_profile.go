package model
import (
	"darkside/db"
	log "github.com/flywithbug/log4go"
	"time"
)



type UserProfileModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	Avatar		string 		`json:"avatar,omitempty" form:"avatar,omitempty"`
	UserId 		string 		`json:"userid,omitempty" form:"userid,omitempty"`
	UpdateTime	int64  		`json:"updatetime,omitempty" form:"updatetime,omitempty"`
}


func (profile *UserProfileModel)UpdateUserProfile()error {
	db := db.DBConf()

	stmt, err := db.Prepare("INSERT profile SET uid =?,userid=?,avatar=?, updatetime=? on duplicate key update uid =?,userid=?,avatar=?, updatetime=?")
	if err != nil{
		log.Warn(err.Error())
		return err
	}
	profile.UpdateTime = time.Now().Unix()
	_, err = stmt.Exec(profile.Uid,profile.UserId,profile.Avatar,profile.UpdateTime,profile.Uid,profile.UserId,profile.Avatar,profile.UpdateTime)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	return err
}



