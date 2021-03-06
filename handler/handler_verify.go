package handler

import (
	"fmt"
	"darkside/common"
	"darkside/config"
	d "darkside/data"
	e "darkside/email"
	"darkside/model"
	"github.com/flywithbug/qcloudsms"
	u "github.com/flywithbug/utils"
	"github.com/gin-gonic/gin"
	"errors"
	"net/http"
	"strings"
	"time"
)

func SendSMSandler(c *gin.Context) {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRespon)
	}()
	tel := model.TelephoneModel{}
	err := c.BindJSON(&tel)
	if err != nil {
		aRespon.SetErrorInfo(d.ErrcodeRequestParamsInvalid, "Param invalid "+err.Error())
		return
	}
	if !common.ValideSMSType(tel.Type) {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara, "Param invalid - verify type not right")
		return
	}
	if !common.ValidePhone(tel.Mobile) {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara, "phone  invalid ")
		return
	}
	err = sendRegisterSMSCode(tel)
	if err != nil {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara, err.Error())
		return
	}
	aRespon.SetSuccessInfo(http.StatusOK, "验证码发送成功")
}

func sendRegisterSMSCode(tel model.TelephoneModel) error {
	sms := model.SMSTXModel{}
	sms.Mobile = tel.Mobile
	sms.Ncode = tel.NCode
	sms.TelModel = tel
	sms.SMStype = tel.Type
	sms.Smscode = strings.ToUpper(u.RandSMSString(6))
	if sms.CheckDidSMSSend() {
		return errors.New("wait 60s")
	}

	if sms.CheckMaxSendSMSCount() {
		return errors.New("tmore than 10 sms has been sended today,")

	}
	//TODO 后续根据不同的type 发送不同的短信模板
	message := fmt.Sprintf("您的验证码是：%s 如非本人操作，请忽略本短信.(http://www.flywithme.top)", sms.Smscode)

	// "您的验证码是：" + sms.Smscode + " 如非本人操作，请忽略本短信.(http://www.flywithme.top)"
	//"欢迎注册案发现场App，请访问http://www.flywithme.top/ 了解更多"
	sms.Messag = message
	sms.Time = time.Now().Unix()
	sms.Status = model.SMSStatusUnChecked
	var err error
	if config.TomlConf().Debug() {
		sms.Errmsg = "debug mock"
		sms.Result = -2
	} else {
		conf := qcloudsms.NewClientConfig()
		conf.AppID = config.TomlConf().Smsc.AppID
		conf.AppKey = config.TomlConf().Smsc.AppKey
		client, err := qcloudsms.NewClient(conf)
		smsReq, err := qcloudsms.SMSService(client)
		ext := qcloudsms.SmsExt{}
		ext.Type = 0
		ext.NationCode = tel.NCode
		resp, err := smsReq.Send(tel.Mobile, sms.Messag, ext)
		sms.Errmsg = resp.ErrMsg
		sms.Sid = resp.Sid
		sms.Result = resp.Result
		sms.Ext = resp.Ext
		sms.Fee = resp.Fee
		if resp.Result != 0 {
			err = errors.New(fmt.Sprintf("%u  %s", resp.Result, resp.ErrMsg))
		}
		if err != nil {
			sms.Status = model.SMSStatusFaild
		}
	}
	sms.InsertSMSInfo()
	return err

}

//邮箱注册暂时关闭
func SendMailHandler(c *gin.Context) {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRespon)
	}()
	mail := model.MailModel{}
	err := c.BindJSON(&mail)
	if err != nil {
		aRespon.SetErrorInfo(http.StatusBadRequest, "Param invalid "+err.Error())
		return
	}
	if !common.ValideMail(mail.Email) {
		aRespon.SetErrorInfo(http.StatusBadRequest, "mail  invalid ")
		return
	}
	err = sendRegistEmailCode(mail)
	if err != nil {
		aRespon.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRespon.SetSuccessInfo(http.StatusOK, "验证码发送成功")

}

func sendRegistEmailCode(mail model.MailModel) error {

	email := model.EmailInfoModel{}
	email.Mail = mail.Email
	email.Type = mail.Type
	email.Verifycode = strings.ToUpper(u.RandSMSString(6))
	email.Status = 1
	email.Message = fmt.Sprintf("您的验证码是：%s 如非本人操作，请忽略本邮件.(http://www.flywithme.top)", email.Verifycode)
	email.Sender = e.USER
	sendmail := e.NewEmail(email.Mail, "案发现场-注册邮件验证", email.Message)
	err := e.SendEmail(sendmail)
	if err != nil {
		email.Status = -1
	}
	email.CreateTime = time.Now().Unix()
	email.InsertSMSInfo()
	return err
}
