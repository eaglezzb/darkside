package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/flywithbug/log4go"
	"sync"
	"net/http"
	"net"
	"encoding/binary"
)

var wg sync.WaitGroup

func Middleware(c *gin.Context)  {
	wg.Add(1)
	go wirteLog(c)
	wg.Wait()
}

func wirteLog(c *gin.Context)  {
	c.Request.ParseForm()
	log.Info("Form:%s,,Path:%s%s",c.Request.Form,RemoteIp(c.Request),c.Request.RequestURI)
	wg.Done()
}

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)


func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}


// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}