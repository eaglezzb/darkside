package main

import (
	"darkside/config"
	"darkside/db"
	"darkside/handler"
	midw "darkside/middleware"
	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
)

func SetLog() {
	w := log.NewFileWriter()
	w.SetPathPattern("./log/log-%Y%M%D.log")
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	log.SetLevel(log.DEBUG)
	log.SetLayout("2006-01-02 15:04:05")
}

func init() {
	config.InitConf("static/config/config.toml")
	db.InitMysql()
	db.InitCache2go()
}

func main() {

	SetLog()
	defer log.Close()

	defer db.DBConf().Close()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	r.Use(midw.Middleware)
	//r.Use(midw.Authorize())
	r.StaticFile("/favicon.ico", "./static/resources/favicon.ico")
	handler.RegisterRouters(r)
	gin.SetMode(config.TomlConf().GinEnv())
	r.Run(config.TomlConf().Server().Listen)
}
