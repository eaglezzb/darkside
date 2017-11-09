package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/flywithbug/darkside/model"
	//"encoding/json"
	log "github.com/flywithbug/log4go"
	"net/http"
	"strconv"
	"regexp"
	"github.com/gin-gonic/gin/json"
	"time"
	_ "fmt"
)

func RegisterHandler(c *gin.Context )  {
	user := m.NewUser()
	json.NewDecoder(c.Request.Body).Decode(&user)
	if !valideUserName(user.UserName) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"用户名不复合要求")
		return
	}
	if !validePassword(user.Password) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"密码不符合要求")
		return
	}
	if !validePhone(user.Phone) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"手机号不符合要求")
		return
	}
	tm := time.Now()
	user.CreateTime = tm.Unix()
	user.UpdateTime = tm.Unix()
	err := user.InsertUser()
	if err != nil {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,err.Error())
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"注册成功",
		"userinfo":user,
	})
}


func LoginHandler(c *gin.Context)  {
	user := m.NewUser()
	json.NewDecoder(c.Request.Body).Decode(&user)

	dbUser ,err := m.FindUserFromDBByName(user.UserName)
	if err != nil  {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"用户名错误")
		return
	}
	if dbUser.Password != user.Password  {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"密码错误")
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK,gin.H{
		"userinfo":dbUser,
		"code":http.StatusOK,
		"message":"",
	})
}


//user/:uid
func GetUserInfoHandler(c *gin.Context)  {
	uid ,_ := strconv.ParseInt(c.Param("uid"),10,64)
	user,err := m.FindUserFromDB(uid)
	if err != nil{
		errCallBack(c,http.StatusOK,http.StatusBadRequest,err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"userinfo":user,
		"code":http.StatusOK,
		"message":"",
	})
}





func valideUserName(name string)bool  {
	reg := regexp.MustCompile("^[a-zA-Z0-9_-]{4,16}$")
	return reg.MatchString(name)
}

func validePassword(name string)bool  {
	reg := regexp.MustCompile("^[a-zA-Z0-9_-]{4,16}$")
	return reg.MatchString(name)
}

func validePhone(name string)bool  {
	reg := regexp.MustCompile("^[1-9][0-9]{4,13}$")
	return reg.MatchString(name)
}

func errCallBack(c *gin.Context,status int,code int,message string)  {
	log.Error(message,c.Request)
	c.JSON(status,gin.H{
		"code":code,
		"message":message,
	})
}
