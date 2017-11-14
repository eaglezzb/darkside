package common

import (
	"regexp"
	"github.com/gin-gonic/gin"
	log "github.com/flywithbug/log4go"
	"strconv"
)



func ValideUserName(name string)bool  {
	reg := regexp.MustCompile("^[a-zA-Z0-9_-]{4,16}$")
	return reg.MatchString(name)
}

func ValidePassword(name string)bool  {
	reg := regexp.MustCompile("^[a-zA-Z0-9_-]{4,16}$")
	return reg.MatchString(name)
}

func ValidePhone(name string)bool  {
	reg := regexp.MustCompile("^[1-9][0-9]{4,13}$")
	return reg.MatchString(name)
}

func ValideSMSType(smstype int)bool  {
	reg := regexp.MustCompile("^[1-2]$")
	return reg.MatchString(strconv.Itoa(smstype))
}


func ValideMail(mail string)bool  {
	reg := regexp.MustCompile("^[A-Za-z0-9_.-\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	return reg.MatchString(mail)
}



func ErrCallBack(c *gin.Context,status int,code int,message string)  {
	log.Error(message,c.Request)
	c.JSON(status,gin.H{
		"code":code,
		"message":message,
	})
}
