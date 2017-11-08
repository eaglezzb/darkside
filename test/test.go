package test

import (
	"github.com/flywithbug/darkside/db"
	"fmt"
	"github.com/flywithbug/darkside/model"
	"time"
	//"github.com/flywithbug/darkside/email"
	_ "github.com/flywithbug/utils"
)

func DBTest()  {

	//dbsqlTest()

	//utils.Test()

	userTest()
	//user,_ := model.FindUserFromDB(20)
	//fmt.Println(user)
	//err := model.DeleteUserFromDB(111)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//fmt.Println("发送邮件")
	//mycontent := "go email"
	//mail := email.NewEmail("369495368@qq.com","test golang email",mycontent)
	//err := email.SendEmail(mail)
	//if err != nil {
	//	fmt.Println(err)
	//}

}

func userTest() {
	user := model.NewUser()
	user.Uid = 139
	user.UserName = "Haryy"
	user.Sex = 1
	user.CreateTime = time.Now().Unix()
	user.Password = "12345sdsd6"
	user.DepartName = "技asdasd"
	fmt.Println(user)
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
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,createtime=?")
	checkErr(err)

	res, err := stmt.Exec("码农", "技术部", time.Now().Unix())
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
		var createtime string
		var age int
		var phone string


		err = rows.Scan(&uid, &username, &department, &createtime,&age,&phone)

		fmt.Println(uid,username,department,createtime,age,phone)
	}
}