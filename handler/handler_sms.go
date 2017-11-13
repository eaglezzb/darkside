package handler

import (
	 "github.com/flywithbug/darkside/model"
	"github.com/flywithbug/qcloudsms"
	"github.com/flywithbug/darkside/config"
	d "github.com/flywithbug/darkside/data"
	u "github.com/flywithbug/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/flywithbug/darkside/common"
	"github.com/kataras/iris/core/errors"
	"fmt"
)

func SendSMSHandler(c *gin.Context )  {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK,aRespon)
	}()
	tel := model.TelephoneModel{}
	err := c.BindJSON(&tel)
	if err != nil{
		aRespon.SetErrorInfo(http.StatusBadRequest,"Param invalid "+err.Error())
		return
	}

	if !common.ValidePhone(tel.Mobile) {
		aRespon.SetErrorInfo(http.StatusBadRequest,"phone  invalid ")
		return
	}
	err = sendRegisterSMSCode(tel)
	if err != nil{
		aRespon.SetErrorInfo(http.StatusBadRequest,err.Error())
		return
	}
	aRespon.SetSuccessInfo(http.StatusOK,"验证码发送成功")
}


func sendRegisterSMSCode(tel model.TelephoneModel)(error)  {
	sms := model.SMSTXModel{}
	sms.Mobile = tel.Mobile
	sms.Ncode = tel.Code
	sms.TelModel = tel
	sms.SMStype = tel.Type
	sms.Smscode = u.RandSMSString(6)
	message := fmt.Sprintf("您的验证码是：%s 如非本人操作，请忽略本短信.(http://www.flywithme.top)",sms.Smscode)

	// "您的验证码是：" + sms.Smscode + " 如非本人操作，请忽略本短信.(http://www.flywithme.top)"
	//"欢迎注册案发现场App，请访问http://www.flywithme.top/ 了解更多"
	sms.Messag = message
	conf := qcloudsms.NewClientConfig()
	conf.AppID = config.TomlConf().Smsc.AppID
	conf.AppKey = config.TomlConf().Smsc.AppKey
	client, err := qcloudsms.NewClient(conf)
	if err != nil {
		return err
	}
	smsReq, err := qcloudsms.SMSService(client)
	//smsReq.
	if err != nil {
		return err
	}
	ext := qcloudsms.SmsExt{}
	ext.Type = 0
	ext.NationCode =tel.Code
	resp, err := smsReq.Send(tel.Mobile, sms.Messag,ext)
	if err != nil{
		return err
	}
	if resp.Result != 0{
		errs := fmt.Sprintf("%u  %s",resp.Result,resp.ErrMsg)
		return errors.New(errs)
	}
	sms.Errmsg = resp.ErrMsg
	sms.Sid = resp.Sid
	sms.Result = resp.Result
	sms.Ext = resp.Ext
	sms.Fee = resp.Fee
	sms.InsertSMSInfo()
	return err

}

