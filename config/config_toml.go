package config

import (
	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"os"
	"github.com/gin-gonic/gin"
)

type  User struct {
	UserName string `toml:"username`
	Password string `toml:"password`
}

type  Server struct {
	Listen string `toml:"listen"`
	CookieCacheTIme int64  `toml:"cookie_cache_time"`
}

type Mysql struct {
	Url       	string 	`toml:"url"`
	DBtype        	string 	`toml:"dbtype"`
}

type  Config struct {
	Env     string 			`toml:"env"`
	User    *User    		`toml:"user"`
	Smsc    	*SMSKey    		`toml:"sms"`
	Servers map[string]*Server 	`toml:"server"`
	Mysqls map[string]*Mysql 	`toml:"mysql"`

}

type SMSKey struct {
	AppID 	string  `toml:"appid"`
	AppKey 	string  `toml:"appkey"`
}

var config *Config

func TomlConf()*Config  {
	return config
}

func InitConf(p string)  {
	if _,err := toml.DecodeFile(p, &config);err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

func (s *Config)Sms()*SMSKey  {
	return s.Smsc
}
func (s *Config)Server()*Server  {
	return s.Servers[s.Env]
}

func (s *Config)Mysql()*Mysql  {
	return s.Mysqls[s.Env]
}

func (s *Config)GinEnv()string  {
	switch  s.Env {
	case "pro":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}




