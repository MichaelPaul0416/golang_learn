package main

import (
	"net/http"
	"web/ch2/db"
	"errors"
	"fmt"
	"html/template"
	"log"
	"strings"
	"os"
)

var logger *log.Logger

func init(){
	file,err := os.OpenFile("chitchat.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil{
		log.Fatalln("failed to open a log file",err)
	}
	logger = log.New(file,"INFO ",log.Ldate|log.Ltime|log.Lshortfile)
}

func session(write http.ResponseWriter,request *http.Request)(s db.Session,err error){
	cookie,err := request.Cookie("_cookie")
	if err == nil{
		s = db.Session{Uuid:cookie.Value}
		//校验session是否存在合法
		if ok,_ := s.CheckValid();!ok{
			err = errors.New("Invalid session")
		}
	}else{
		//err not nil,redirect to login
		info("empty cookie and redirect to login")
		http.Redirect(write,request,"/login",302)
	}
	return
}


func generateHTML(write http.ResponseWriter,data interface{},fileNames... string){
	//文件的路径都是相对src的上一层来说的
	var files []string
	for _,file := range fileNames{
		files = append(files,fmt.Sprintf("src/web/ch2/templates/%s.html",file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(write,"layout",data)
}

func parseTemplateFiles(fileNames... string)(t *template.Template){
	var files []string
	t = template.New("layout")
	for _,file := range fileNames{
		files = append(files,fmt.Sprintf("src/web/ch2/templates/%s.html",file))
	}

	templates := template.Must(t.ParseFiles(files...))
	return templates
}

//log warning
func warning(args... interface{}){
	logger.SetPrefix("WARNING")
	logger.Println(args)
	log.SetPrefix("WARNING")
	log.Println(args)
}

func danger(args...interface{}){
	logger.SetPrefix("ERROR")
	logger.Println(args)
	log.SetPrefix("ERROR")
	log.Println(args)
}

func info(args...interface{}){
	logger.SetPrefix("INFO")
	logger.Println(args)
	log.SetPrefix("INFO")
	log.Println(args)
}

func error_message(w http.ResponseWriter,r *http.Request,msg string){
	url := []string{"/err?msg=",msg}
	//string.Join类似于StringBuilder
	http.Redirect(w,r,strings.Join(url,""),302)
}