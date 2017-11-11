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
		aRespon.SetErrorInfo(http.StatusBadRequest,"Param invalid "+err.Error())
		return
	}

	if !common.ValideUserName(user.UserName) {
		aRespon.SetErrorInfo(http.StatusBadRequest,"username invalid")
		return
	}
	if !common.ValidePassword(user.Password) {
		aRespon.SetErrorInfo(http.StatusBadRequest,"password invalid ")
		return
	}
	if !common.ValidePhone(user.Phone) {
		aRespon.SetErrorInfo(http.StatusBadRequest,"phone  invalid ")
		return
	}
	tm := time.Now()
	user.CreateTime = tm.Unix()
	user.UpdateTime = tm.Unix()
	err = user.InsertUser()
	if err != nil {
		aRespon.SetErrorInfo(http.StatusBadRequest,"db insert faild "+err.Error())
		return
	}
	aRespon.AddResponseInfo("user",user)
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
		aRespon.SetErrorInfo(http.StatusBadRequest, "Params invalid " + err.Error())
		return
	}
	dbUser, err := m.CheckUserNameAndPass(user.UserName, user.Password)
	if err != nil {
		aRespon.SetErrorInfo(http.StatusBadRequest, err.Error())
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
		aRespon.SetErrorInfo(http.StatusBadRequest,err.Error())
		return
	}
	user.Password = ""
	user.OldPassword = ""
	aRespon.AddResponseInfo("user",user)
}

