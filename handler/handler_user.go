package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/flywithbug/darkside/model"
	 "github.com/flywithbug/darkside/common"
	"github.com/gin-gonic/gin/json"
	d "github.com/flywithbug/darkside/data"
	"time"
	_ "fmt"
	"reflect"
	"fmt"
	"net/http"
	"strconv"
)

func RegisterHandler(c *gin.Context )  {
	user := m.NewUser()
	aRespon := d.NewResponse()
	defer func() {
		c.JSON(http.StatusOK,aRespon)
	}()

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


func LoginHandler(c *gin.Context)  {
	user := m.NewUser()
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil  {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,"数据解析错误")
		return
	}
	dbUser ,err := m.FindUserFromDBByName(user.UserName)
	if err != nil  {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,"用户名错误")
		return
	}
	fmt.Println(dbUser.Password, user.Password)
	if dbUser.Password != user.Password  {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,"密码错误")
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"userinfo":dbUser,
		"code":http.StatusOK,
		"message":"",
	})
}



func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}


//user/:uid
func GetUserInfoHandler(c *gin.Context)  {
	uid ,_ := strconv.ParseInt(c.Param("uid"),10,64)
	user,err := m.FindUserFromDB(uid)
	jsons,_ := json.Marshal(user)
	fmt.Println(string(jsons))


	if err != nil{
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,err.Error())
		return
	}
	user.Password = ""
	user.OldPassword = ""
	c.JSON(http.StatusOK,gin.H{
		"userinfo":user,
		"code":http.StatusOK,
		"message":"",
	})
}

