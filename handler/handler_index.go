package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/flywithbug/darkside/test"
)



//golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
func Index(ctx *gin.Context)  {

	test.Test()

	ctx.JSON(200,gin.H{
		"errno":"0",
		"msg":"Index.html",
	})
}
