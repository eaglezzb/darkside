package middleware

//import (
//	"github.com/gin-gonic/gin"
//	"regexp"
//	"fmt"
//)
//
//
//
//func Authorize() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		authHeader := c.Request.Header.Get("Authorization")
//		fmt.Println("authHeader", authHeader)
//		r, _ := regexp.Compile("^Bearer (.+)$")
//
//		match := r.FindStringSubmatch(authHeader)
//
//		if len(match) == 0 {
//			c.AbortWithStatus(401)
//			return
//		}
//		tokenString := match[1]
//
//		if len(tokenString) == 0 {
//			c.AbortWithStatus(401)
//			return
//		}
//
//		user := "tokenString"
//
//		c.Set("user", user)
//		c.Set("userID", user)
//		c.Next()
//	}
//}