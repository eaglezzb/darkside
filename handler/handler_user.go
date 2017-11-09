package handler

import (
	"github.com/gin-gonic/gin"
	m "github.com/flywithbug/darkside/model"
	"net/http"
	"strconv"
	"github.com/flywithbug/darkside/common"
	"github.com/gin-gonic/gin/json"
	"time"
	_ "fmt"
	"reflect"
	"fmt"
)

func RegisterHandler(c *gin.Context )  {
	user := m.NewUser()
	json.NewDecoder(c.Request.Body).Decode(&user)
	if !common.ValideUserName(user.UserName) {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,"用户名不复合要求")
		return
	}
	if !common.ValidePassword(user.Password) {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,"密码不符合要求")
		return
	}
	if !common.ValidePhone(user.Phone) {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,"手机号不符合要求")
		return
	}
	tm := time.Now()
	user.CreateTime = tm.Unix()
	user.UpdateTime = tm.Unix()
	err := user.InsertUser()
	if err != nil {
		common.ErrCallBack(c,http.StatusOK,http.StatusBadRequest,err.Error())
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
	c.JSON(http.StatusOK,gin.H{
		"userinfo":user,
		"code":http.StatusOK,
		"message":"",
	})
}

