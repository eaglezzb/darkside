package model

import (
	"github.com/flywithbug/darkside/db"
	log "github.com/flywithbug/log4go"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"time"
)
const (
	SMSStatusFaild           	= -1 //短信发送失败
	SMSStatusUnChecked           	= 1 //未校验
	SMSStatusChecked           	= 2 //已校验
	SMSStatusOverTime           	= 3 //校验时超过有效时间


	SMSTypeRegister    		= 1 //注册
	SMSTypeChangePassword		= 2 //修改密码


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

type TelephoneModel struct {
	Code		string 			`json:"code,omitempty" form:"ncode,omitempty"`
	Mobile		string 			`json:"mobile,omitempty" form:"mobile,omitempty"`
	Type		int 			`json:"type,omitempty" form:"type,omitempty"` //1 用户注册类型
}

func (sms *SMSTXModel)InsertSMSInfo()error {
	fmt.Println(sms)
	db := db.DBConf()
	stmt, err := db.Prepare("INSERT smstx SET type=?,message=?,result=?,time=?,ext=?,mobile=?,ncode=?,errmsg=?,sid=?,fee=?,smscode=?,status=?")
	checkSMSErr(err)
	_, err = stmt.Exec(sms.SMStype,sms.Messag,sms.Result,sms.Time,sms.Ext,sms.Mobile,sms.Ncode,sms.Errmsg,sms.Sid,sms.Fee,sms.Smscode,sms.Status)
	if err != nil {
		log.Warn(err.Error())
	}
	//手机数据短信send次数自增
	stmt, err = db.Prepare("INSERT telephone SET mobile=?,ncode=? on duplicate key update scount=scount+1,ncode=?")
	checkSMSErr(err)
	_, err1 := stmt.Exec(sms.Mobile,sms.Ncode,sms.Ncode)
	if err1 != nil {
		log.Warn(err1.Error())
	}
	return err
}

func (sms *SMSTXModel)MarkSmsVerifyCode(status int)error  {
	db := db.DBConf()
	stmt,err := db.Prepare("update smstx set status=? where uid=?")
	checkSMSErr(err)
	if err != nil{
		return err
	}
	_,err = stmt.Exec(status,sms.Uid)
	checkSMSErr(err)
	return err
}

func CheckPhoneAndVerifyCode(phone string,verifycode string)(SMSTXModel,error)  {
	var sms SMSTXModel
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, mobile, time, smscode, status,type FROM smstx WHERE mobile=? and smscode=?", phone,verifycode).
			Scan(&sms.Uid,
			&sms.Mobile, &sms.Time,&sms.Smscode,&sms.Status,&sms.SMStype)

	checkSMSErr(err)
	if err != nil {
		return sms,err
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
//todo 一天内只能发送10次验证码
func (sms *SMSTXModel)CheckDidSMSSend()bool  {
	db := db.DBConf()
	err := db.QueryRow("SELECT uid, mobile, time, smscode, status,type FROM smstx WHERE mobile=? and time>?", sms.Mobile,time.Now().Unix()-60).
		Scan(&sms.Uid,
		&sms.Mobile, &sms.Time,&sms.Smscode,&sms.Status,&sms.SMStype)
	checkSMSErr(err)
	if err != nil{
		return false
	}
	return true
}






//func CheckUserNameAndPass(name string,pass string)(UserInfoModel,error)  {
//	var user UserInfoModel
//	db := db.DBConf()
//	err := db.QueryRow("SELECT uid, username, departname, password, sex, userid, phone, phoneprefix, createtime, updatetime, state, authtoken, mail, oldpassword FROM userinfo WHERE username=?", name).
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


func checkSMSErr(err error) {
	if err != nil {
		log.Warn(err.Error())
	}
}