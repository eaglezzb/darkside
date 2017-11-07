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
	nameValid := m.CheckUserNameValid(user.UserName)
	if nameValid == false {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"该用户名已被注册")
		return
	}

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
	phonePrefixstr := c.PostForm("phoneprefix")
	user.PhonePrefix ,_ = strconv.ParseInt(phonePrefixstr,10,64)

	phonestr := c.PostForm("phone")
	if !validePhone(phonestr) {
		errCallBack(c,http.StatusOK,http.StatusBadRequest,"手机号不符合要求")
		return
	}
	user.Phone ,_ = strconv.ParseInt(phonestr,10,64)
	tm := time.Now()
	user.CreateTime = tm.Unix()
	user.UpdateTime = tm.Unix()
	err := user.InsertUser()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusBadRequest,
			"message":"用户生成失败",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"message":"注册成功",
		"userinfo":user,
	})
}

func errCallBack(c *gin.Context,status int,code int,message string)  {
	c.JSON(status,gin.H{
		"code":code,
		"message":message,
	})
}



func LoginHandler(c *gin.Context)  {
	//user := m.NewUser()




}


//user/:uid
func GetUserInfoHandler(c *gin.Context)  {
	uid ,_ := strconv.ParseInt(c.Param("uid"),10,64)
	user,err := m.FindUserFromDB(uid)
	if err != nil{
		log.Error(err.Error(),err)
		c.JSON(http.StatusOK,gin.H{
			"userinfo":nil,
			"code":404,
			"message":"未查询到该用户信息",
		})
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