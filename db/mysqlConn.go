package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/flywithbug/darkside/config"
	"fmt"
)

var db *sql.DB
var err error

func DBConf()*sql.DB  {
	return db
}
func InitMysql()  {
	fmt.Println(config.TomlConf().Mysql().DBtype, config.TomlConf().Mysql().Url)
	db, err = sql.Open(config.TomlConf().Mysql().DBtype, config.TomlConf().Mysql().Url)
	db.SetMaxIdleConns(200)
	db.SetMaxOpenConns(100)
	checkErr(err)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}






