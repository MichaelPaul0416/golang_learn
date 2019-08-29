package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func writeExample(w http.ResponseWriter,r *http.Request){
	//如果没有手动指定返回的消息类别，那么golang将会按照返回数据的前512字节进行判断
	str := `<html><head><title>Go Web Programming</title></head><body><h1>Hello World</h1></body></html>`
	w.Write([]byte(str))
}

func writePlain(w http.ResponseWriter,r *http.Request){
	str := `{"launchedChannel":"2","data":[{"phoneNo":"","creditNo":"310221197406102822","creditType":"00","creditName":"杨济如","accountNo":"5240112130010111","certificateInvestor":"DS"}`
	w.Write([]byte(str))
}

func writeRedirect(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Location","https://www.baidu.com")
	//接受一个整数作为http的响应码，这个方法并不是设置响应头，在调用这个方法之后，用户可以继续对ResponseWriter进行写入
	//但是不能对响应的首部做任何写入的操作
	//如果用户在调用Write方法之前没有执行过WriteHeader方法，那么默认就会使用200作为响应码
	w.WriteHeader(302)//写了响应码之后就不能在对首部Header进行设置了
}

type bean struct {
	name string
	id int
	parent string
}
func writeJson(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	post := &bean{
		name:"Ioc",
		id:1,
		parent:"Spring",
	}

	str,err := json.Marshal(post)
	if err != nil{
		fmt.Printf("error:%v\n",err)
	}else{
		fmt.Printf("json:%s\n",string(str))
	}
	w.Write(str)
}

func main(){
	server := http.Server{
		Addr:"0.0.0.0:8080",
	}

	http.HandleFunc("/html/",writeExample)
	http.HandleFunc("/plain/",writePlain)
	http.HandleFunc("/redirect/",writeRedirect)
	http.HandleFunc("/json/",writeJson)
	server.ListenAndServe()
}
