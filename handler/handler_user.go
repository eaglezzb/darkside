package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/brasbug/darkside/model"
	//"encoding/json"
	log "github.com/brasbug/log4go"
	"net/http"
	"strconv"
	"regexp"
	"fmt"
)

func RegisterHandler(c *gin.Context )  {
	user := m.NewUser()
	user.UserName = c.PostForm("username")
	if valideUserName(user.UserName) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusBadRequest,
			"message":"用户名不复合要求",
		})
		return
	}
	user.Password = c.PostForm("password")
	if valideUserName(user.Password) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusBadRequest,
			"message":"密码不符合要求",
		})
		return
	}
	user.Sex,_ = strconv.Atoi(c.PostForm("sex"))
	user.DepartName = c.PostForm("departname")
	user.PhonePrefix ,_ = strconv.ParseInt(c.PostForm("phoneprefix"),10,6)
	user.Phone ,_ = strconv.ParseInt(c.PostForm("phone"),10,11)

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

func valideUserName(name string)bool  {
	reg := regexp.MustCompile("/^[a-zA-Z0-9_-]{4,16}$/")
	return reg.MatchString(name)
}

func validePassword(name string)bool  {
	reg := regexp.MustCompile("/^[a-zA-Z0-9_-]{4,16}$/")
	return reg.MatchString(name)
}


//user/:uid
func GetUserInfo(c *gin.Context)  {
	uid ,_ := strconv.ParseInt(c.Param("uid"),10,64)
	user,err := m.FindUserFromDB(uid)
	fmt.Println(user)
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

