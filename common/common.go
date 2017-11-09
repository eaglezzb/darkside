package common

import (
	"regexp"
	"github.com/gin-gonic/gin"
	log "github.com/brasbug/log4go"
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



func ErrCallBack(c *gin.Context,status int,code int,message string)  {
	log.Error(message,c.Request)
	c.JSON(status,gin.H{
		"code":code,
		"message":message,
	})
}
