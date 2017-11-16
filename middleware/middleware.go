package middleware

import (
	"github.com/gin-gonic/gin"
	//log "github.com/flywithbug/log4go"
	"sync"
	//"github.com/flywithbug/utils"
)

var wg sync.WaitGroup

func Middleware(c *gin.Context)  {
	wg.Add(1)
	go wirteLog(c)
	wg.Wait()
}

func wirteLog(c *gin.Context)  {
	//buf := make([]byte,1024)
	//n, err := c.Request.Body.Read(buf)
	//log.Info("Form:%s,,ip:%s path:%s err:%s",string(buf[0:n]),utils.RemoteIp(c.Request),c.Request.RequestURI,err.Error())
	wg.Done()
}

