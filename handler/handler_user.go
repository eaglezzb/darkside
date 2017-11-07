package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/brasbug/darkside/model"
	//"encoding/json"
	log "github.com/brasbug/log4go"
	"net/http"
	"strconv"
	"regexp"
	"time"
)

func RegisterHandler(c *gin.Context )  {
	user := m.NewUser()
	user.UserName = c.PostForm("username")

	if !valideUserName(user.UserName) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"用户名不复合要求")
		return
	}

	user.Password = c.PostForm("password")
	if !validePassword(user.Password) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"密码不符合要求")
		return
	}
	user.Sex,_ = strconv.Atoi(c.PostForm("sex"))
	user.DepartName = c.PostForm("departname")
	user.PhonePrefix = c.PostForm("phoneprefix")
	user.Phone  = c.PostForm("phone")
	if !validePhone(c.PostForm("phone")) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"手机号不符合要求")
		return
	}
	user.Phone  = c.PostForm("phone")
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
	username := c.PostForm("username")
	password := c.PostForm("password")
	user ,err := m.FindUserFromDBByName(username)
	if err != nil  {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"用户名错误")
		return
	}
	if user.Password != password  {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"密码错误")
		return
	}


	c.JSON(http.StatusOK,gin.H{
		"userinfo":user,
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
