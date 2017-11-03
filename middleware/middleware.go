package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/brasbug/log4go"
	"sync"
)

var wg sync.WaitGroup

func Middleware(c *gin.Context)  {
	wg.Add(1)
	go wirteLog(c)
	wg.Wait()
}

func wirteLog(c *gin.Context)  {
	c.Request.ParseForm()
	log.Info("Form:%s,Path:%s%s",c.Request.Form,c.Request.Host,c.Request.RequestURI)
	wg.Done()
}