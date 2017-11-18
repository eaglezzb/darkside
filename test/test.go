package test

import (
	"github.com/flywithbug/darkside/db"
	"github.com/flywithbug/darkside/model"
	"time"
	"github.com/flywithbug/darkside/email"
	_ "github.com/flywithbug/utils"
	_ "github.com/flywithbug/darkside/config"
	"github.com/flywithbug/darkside/config"
	"github.com/flywithbug/utils"

	"fmt"
)

func Test()  {

	//dbsqlTest()

	//utils.Test()

	//userTest()
	//user,_ := model.FindUserFromDB(1000048)
	//fmt.Println(user)
	//err := model.DeleteUserFromDB(111)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//sendMail()
	//fmt.Println(utils.ConvertInt2String(user))
	//
	//fmt.Println(utils.ConvertInt2String(2002002))
	//fmt.Println(utils.ConvertInt2String("232aass"))
	//var n int64
	//n = 122111111
	//fmt.Println(utils.ConvertInt2String(n))
	insertMail()

}

func insertLocation()  {
	location := model.LocationModel{}
	location.Uid = 10000048
	location.Longitude = 12222322
	location.Latitude = 22999499
	location.UpdateLocation()

}

func insertMail()  {
	mail := model.EmailInfoModel{}

	mail.Mail = "2323@qq.com"
	mail.Verifycode = "34la24"
	mail.Message = "请查收验证码：12322"
	mail.Type = 1
	mail.Status = -1
	mail.CreateTime = time.Now().Unix()
	mail.Sender = email.USER
	mail.InsertSMSInfo()


}

func sendMail()  {
	fmt.Println("发送邮件")

	mycontent := "请查收验证码：12322"
	mail := email.NewEmail("myworldmine@163.com","案发现场-注册邮件验证",mycontent)
	err := email.SendEmail(mail)
	fmt.Println(err)
}

func randomString()  {
	fmt.Println(utils.RandSMSString(6))
}

func testPhoneDB()  {
	fmt.Println(config.TomlConf().Smsc)

	var telModel model.SMSTXModel
	telModel.Messag = "abdadasd"
	telModel.SMStype = 1;
	telModel.Time = time.Now().Unix()
	telModel.Mobile = "17602198928"
	telModel.Smscode = "adb234"
	telModel.TelModel.NCode = "86"
	telModel.TelModel.Mobile = telModel.Mobile
	telModel.Result = 2
	telModel.Ncode = "86"
	telModel.Fee = 1
	fmt.Println(telModel.InsertSMSInfo())
}

func userTest() {
	user := model.NewUser()
	user.Uid = 139
	user.UserName = "Haryy"
	user.Sex = 1
	user.CreateTime = time.Now().Unix()
	user.Password = "12345sdsd6"

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
	stmt, err := db.Prepare("INSERT user SET username=?,departname=?,createtime=?")
	checkErr(err)

	res, err := stmt.Exec("码农", "技术部", time.Now().Unix())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update user set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("码农二代", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM user")
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