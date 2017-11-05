package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/brasbug/darkside/model"
	//"encoding/json"
	log "github.com/brasbug/log4go"
	"net/http"
	"strconv"
)

func RegisterHandler(ctx *gin.Context )  {



}

//user/:uid
func GetUserInfo(c *gin.Context)  {
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

