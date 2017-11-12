package handler

import (
	"fmt"
	"github.com/flywithbug/darkside/model"
	"github.com/flywithbug/qcloudsms"
	"github.com/flywithbug/darkside/config"
	u "github.com/flywithbug/utils"
	"github.com/gin-gonic/gin"
)

func SendSMSHandler(c *gin.Context )  {


}



func SendRegisterSMSCode(tel model.TelephoneModel)(string,error)  {

	sms := model.SMSTXModel{}
	sms.Mobile = tel.Mobile
	sms.Ncode = tel.Code
	sms.TelModel = tel
	sms.SMStype = "0"
	sms.Smscode = u.RandSMSString(6)
	sms.Messag = fmt.Sprintf("您的验证码是:%s,请于5分钟内填写,如非本人操作,请忽略本短信。 更多请访问网站我们的官网 http://www.flywithme.top", sms.Smscode)
	conf := qcloudsms.NewClientConfig()
	conf.AppID = config.TomlConf().Smsc.AppID
	conf.AppKey = config.TomlConf().Smsc.AppKey
	client, err := qcloudsms.NewClient(conf)
	if err != nil {
		return "",err
	}
	smsReq, err := qcloudsms.SMSService(client)
	if err != nil {
		return "",err
	}
	ext := qcloudsms.SmsExt{}
	ext.Type = 0
	ext.NationCode =tel.Code
	fmt.Println(sms)

	resp, err := smsReq.Send(tel.Mobile, "欢迎注册案发现场App，请访问http://www.flywithme.top/ 了解更多",ext)
	if err != nil{
		return "",err
	}
	sms.Errmsg = resp.ErrMsg
	sms.Sid = resp.Sid
	sms.Result = resp.Result
	sms.Ext = resp.Ext
	sms.Fee = resp.Fee
	fmt.Println(sms)
	sms.InsertSMSInfo()
	return sms.Smscode,err

}

