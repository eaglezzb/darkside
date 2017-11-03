package handler

import (
	"github.com/gin-gonic/gin"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"fmt"
	"html/template"
	"os"
)

func UploadFile(c *gin.Context)  {

	c.Request.ParseMultipartForm(32<< 20)
	file ,handler,err := c.Request.FormFile("uploadfile")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer  file.Close()
	fmt.Fprintf(c.Writer,"%v",handler.Header)
	f,err := os.OpenFile("./uploadfile/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f,file)

}

func UploadIndex(c *gin.Context)  {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h,strconv.FormatInt(crutime,10))
	token := fmt.Sprintf("%x",h.Sum(nil))
	t,_:= template.ParseFiles("static/gtpl/upload.gtpl")
	t.Execute(c.Writer,token)
}
