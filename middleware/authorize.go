package middleware

import (
	//"github.com/gin-gonic/gin"
	//"regexp"
	//"fmt"
	//"github.com/flywithme/darkside/model"
)


//func Authorize()gin.HandlerFunc  {
	//return func(c *gin.Context) {
	//	authHeader := c.Request.Header.Get("Authorization")
	//	r,_:= regexp.Compile("^Bearer (.+)$")
	//
	//	match := r.FindStringSubmatch(authHeader)
	//
	//	if len(match) == 0 {
	//		c.AbortWithStatus(401)
	//		return
	//	}
	//
	//	tokenString := match[1]
	//	if len(tokenString) == 0 {
	//		c.AbortWithStatus(401)
	//		return
	//	}
	//
	//	fmt.Println(tokenString)
	//	user := model.NewUser()
	//	user.Password = "abc"
	//	user.Uid ="1"
	//	user.Name = "a"
	//
	//	c.Set("user",user)
	//	c.Next()
	//
	//
	//}


//}