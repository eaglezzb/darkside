package handler



import (
	"github.com/gin-gonic/gin"

	"strings"
	"fmt"
)


var routers = map[string]gin.HandlerFunc{
	"GET      /":Index,
	"GET      /upload":UploadIndexHandler,
	"POST     /uploadfile":UploadFileHandler,

	"GET      /user/:uid":GetUserInfoHandler,
	"POST     /user/registeruser":RegisterHandler,
	"POST    /user/login":LoginHandler,
}


func RegisterRouters(r *gin.Engine)  {
	for route, handler := range routers {
		route = strings.TrimSpace(route)
		fmt.Println(route)
		method := "GET"
		//空格
		idx := strings.Index(route, " ")

		if idx > -1 {
			method = route[:idx]
			route = strings.TrimSpace(route[idx+1:])
		}
		switch method {
		case "POST":
			r.POST(route, handler)
		case "PUT":
			r.PUT(route, handler)
		case "HEAD":
			r.HEAD(route, handler)
		case "DELETE":
			r.DELETE(route, handler)
		case "OPTIONS":
			r.OPTIONS(route, handler)
		default:
			r.GET(route, handler)
		}
	}

}

