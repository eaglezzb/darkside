package test

import (
	"github.com/brasbug/darkside/db"
	"fmt"
	"github.com/brasbug/darkside/model"
	"time"
)

func DBTest()  {

	//dbsqlTest()

	insertTest()
}

func insertTest() {

	user := model.NewUser()
	user.UserName = "Jack"
	//user.Sex = 1
	user.CreateTime = time.Now()
	user.Password = "123456"
	user.DepartName = "技术部"
	user.InsertUser()
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func dbsqlTest()  {

	var db = db.DBConf()

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("码农", "技术部", "2017-11-06")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("码农二代", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		var age int
		var phone string


		err = rows.Scan(&uid, &username, &department, &created,&age,&phone)

		fmt.Println(uid,username,department,created,age,phone)
	}
}