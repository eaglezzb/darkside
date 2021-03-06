package model

import (
	"darkside/db"
	log "github.com/flywithbug/log4go"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"time"
	"database/sql"
)
const (
	SMSStatusFaild           	= -1 //短信发送失败
	SMSStatusUnChecked           	= 1 //未校验
	SMSStatusChecked           	= 2 //已校验
	SMSStatusOverTime           	= 3 //校验时超过有效时间


	SMSTypeRegister    		= 1 //注册
	SMSTypeChangePassword		= 2 //修改密码



	//CallBackCode






)



type SMSTXModel struct {
	Uid  		int64 		`json:"uid,omitempty" form:"uid,omitempty"`
	SMStype 	int 		`json:"type,omitempty" form:"type,omitempty"`   //1 用户注册类型
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
	Status          int 		`json:"status,omitempty" form:"status,omitempty"`
	TelModel	TelephoneModel	`json:"tel,omitempty" form:"tel,omitempty"`
}



func (sms *SMSTXModel)InsertSMSInfo()error {
	fmt.Println(sms)
	db := db.DBConf()
	stmt, err := db.Prepare("INSERT smstx SET type=?,message=?,result=?,time=?,ext=?,mobile=?,ncode=?,errmsg=?,sid=?,fee=?,smscode=?,status=?")
	if err != nil{
		log.Warn(err.Error())
		return err
	}
	_, err = stmt.Exec(sms.SMStype,sms.Messag,sms.Result,sms.Time,sms.Ext,sms.Mobile,sms.Ncode,sms.Errmsg,sms.Sid,sms.Fee,sms.Smscode,sms.Status)
	if err != nil {
		log.Warn(err.Error())
		return err
	}
	var tel TelephoneModel
	tel.Mobile = sms.Mobile
	tel.NCode = sms.Ncode
	return tel.UpdateSendCount2DB()
}

func (sms *SMSTXModel)MarkSmsVerifyCode(status int)error  {
	db := db.DBConf()
	stmt,err := db.Prepare("update smstx set status=? where uid=?")
	if err != nil{
		log.Warn(err.Error())
		return err
	}
	_,err = stmt.Exec(status,sms.Uid)
	if err != nil {
		log.Warn(err.Error())
	}
	return err
}

func CheckPhoneAndVerifyCode(phone string,verifycode string)(SMSTXModel,error)  {
	var sms SMSTXModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, mobile, time, smscode, status,type FROM smstx WHERE mobile=? and smscode=?", phone,verifycode).
			Scan(&sms.Uid,
			&sms.Mobile, &sms.Time,&sms.Smscode,&sms.Status,&sms.SMStype)

	if err != nil {
		log.Warn(err.Error())
		return sms,errors.New("invalid code")
	}
	if sms.Status != SMSStatusUnChecked {
		return sms, errors.New("invalid code")
	}
	
	if time.Now().Unix() - sms.Time  > 1800{
		sms.MarkSmsVerifyCode(SMSStatusOverTime)
		return sms,errors.New("verify time out")
	}
	return sms,nil
}


//判断60秒内是否发送过短信验证码
func (sms *SMSTXModel)CheckDidSMSSend()bool  {
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, mobile, time, smscode, status,type FROM smstx WHERE mobile=? and time>?", sms.Mobile,time.Now().Unix()-60).
		Scan(&sms.Uid,
		&sms.Mobile, &sms.Time,&sms.Smscode,&sms.Status,&sms.SMStype)
	if err != nil {
		return false
	}
	return true
}

//服务方限制 每日10条
func (sms *SMSTXModel)CheckMaxSendSMSCount()bool  {
	db := db.DBConf()
	rows, err := db.Query("SELECT COUNT(*) as count FROM smstx WHERE mobile=? and time>?", sms.Mobile,time.Now().Unix()-60*60*24)
	if err != nil {
		log.Warn(err.Error())
	}
	if checkCount(rows) > 9 {
		return true
	}
	return false
}


func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err:= rows.Scan(&count)
		if err != nil {
			log.Warn(err.Error())
		}
	}
	rows.Close()
	return count
}





//func CheckUserNameAndPass(name string,pass string)(UserModel,error)  {
//	var user UserModel
//	db := db.DBConf()
//	err := db.QueryRow("SELECT uid, username, departname, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM user WHERE username=?", name).
//		Scan(&user.Uid,
//		&user.UserName, &user.DepartName, &user.Password, &user.Sex, &user.UserId, &user.Phone, &user.PhonePrefix,
//		&user.CreateTime, &user.UpdateTime,&user.State,&user.Authtoken,&user.Mail,&user.OldPassword)
//
//	if user.Password != pass {
//		err = errors.New("password not right")
//		user = NewUser()
//	}
//	user.Password = ""
//	user.OldPassword = ""
//	checkErr(err)
//	return user,err
//}


//func checkSMSErr(err error) {
//	if err != nil {
//		log.Warn(err.Error())
//	}
//}