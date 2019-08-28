package main

import (
	"net/http"
	"fmt"
	"time"
	"runtime"
	"reflect"
)

type MyHandler struct{}

//实现了ServeHTTP接口方法的，都是属于处理器
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

type NameHandler struct{}

func (h *NameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "golang http server")
}

//指定一个处理器函数
func timeNow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "time now:%s\n", DateTimeNow())
}

func DateTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func log(h http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("handler function called-" + name)
		h(writer,request)
	}
}

func main() {
	handler := MyHandler{}
	//server := http.Server{
	//	Addr:    "127.0.0.1:8080",
	//一个实现了ServeHTTP的Handler处理所有的请求
	//Handler: &handler,
	//}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	nameHandler := NameHandler{}
	http.Handle("/wel", &handler)
	http.Handle("/info", &nameHandler)
	//HandleFunc 函数将timeNow函数转换为一个Handler，将它与DefaultServeMux绑定，以此简化创建并且绑定Handler的工作
	http.HandleFunc("/time", log(timeNow))

	server.ListenAndServe()
}
