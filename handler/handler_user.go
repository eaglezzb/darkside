package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/flywithbug/darkside/model"
	 "github.com/flywithbug/darkside/common"
	d "github.com/flywithbug/darkside/data"
	"time"
	_ "fmt"
	"fmt"
	"net/http"
	"strconv"
)

func RegisterHandler(c *gin.Context )  {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK,aRespon)
	}()
	user := m.NewUser()
	err := c.BindJSON(&user)
	fmt.Println(user,err)
	if err != nil{
		aRespon.SetErrorInfo(d.ErrcodeRequestParamsInvalid,"Param invalid "+err.Error())
		return
	}
	//if !common.ValideUserName(user.UserName) {
	//	aRespon.SetErrorInfo(http.StatusBadRequest,"username invalid")
	//	return
	//}
	if !common.ValidePassword(user.Password) {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,"password invalid ")
		return
	}
	if !common.ValidePhone(user.Phone) {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,"phone  invalid ")
		return
	}

	sms ,err := m.CheckPhoneAndVerifyCode(user.Phone,user.VerifyCode)
	if err != nil {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,err.Error())
		return
	}
	if sms.SMStype != m.SMSTypeRegister {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,"Incorrect type of verification code")
		return
	}

	tm := time.Now()
	user.CreateTime = tm.Unix()
	user.UpdateTime = tm.Unix()
	err = user.InsertUser()
	if err != nil {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,err.Error())
		return
	}
	user.VerifyCode = ""
	user.Password = ""
	aRespon.AddResponseInfo("code",http.StatusOK)
	aRespon.AddResponseInfo("user",user)
	sms.MarkSmsVerifyCode(m.SMSStatusChecked)


}


func LoginHandler(c *gin.Context) {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRespon)
	}()
	user := m.NewUser()
	err := c.BindJSON(&user)
	fmt.Println(err,user)
	if err != nil {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara, "Params invalid " + err.Error())
		return
	}
	dbUser, err := m.UserLogin(user.UserName, user.Password)
	if err != nil {
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara, err.Error())
		return
	}
	aRespon.AddResponseInfo("user", dbUser)
}


//user/:uid
func GetUserInfoHandler(c *gin.Context)  {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRespon)
	}()
	uid ,_ := strconv.ParseInt(c.Param("uid"),10,64)
	user,err := m.FindUserFromDB(uid)
	if err != nil{
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,err.Error())
		return
	}
	aRespon.AddResponseInfo("user",user)
}

//user/:userid
func GetUserInfoUserIdHandler(c *gin.Context)  {
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRespon)
	}()
	userid := c.Param("userid")
	fmt.Println("userid",userid)
	user,err := m.FindUserFromDBByUserid(userid)
	if err != nil{
		aRespon.SetErrorInfo(d.ErrCodeRequestInvalidPara,err.Error())
		return
	}
	aRespon.AddResponseInfo("user",user)
}

