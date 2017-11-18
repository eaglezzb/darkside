package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/flywithbug/darkside/data"
	"net/http"
	"github.com/flywithbug/darkside/model"
	"github.com/flywithbug/log4go"
	"fmt"
)




func UpdateLocation(c *gin.Context)  {
	aRespon := data.NewResponse()
	defer func() {
		c.JSON(http.StatusOK,aRespon)
	}()
	location := model.LocationModel{}
	fmt.Println(location)
	err := c.BindJSON(&location)
	if err != nil{
		log4go.Error(err.Error())
		aRespon.SetErrorInfo(data.ErrcodeRequestParamsInvalid,"Param invalid "+err.Error())
		return
	}

}