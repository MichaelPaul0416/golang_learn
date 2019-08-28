package main

import (
	"net/http"
	"fmt"
	"time"
	"runtime"
	"reflect"
	"bytes"
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

// 新增函数作为入参传入的demo
type inner struct {
	name string
	age int
	desc string
}

func (i *inner) String() string{
	var buf bytes.Buffer
	fmt.Fprintf(&buf,"inner{name:%s ,age:%d ,desc:%s}\n",i.name,i.age,i.desc)
	return buf.String()
}
//定义一个函数类型
type innerFunc func(name string,age int)(in inner)

func combineFunc(pre string,f innerFunc){
	i := f("Paul",10)
	fmt.Printf("info:%s->%s\n",pre,&i)
}

//接受函数做为入参,同时出参也是一个函数
func log(h http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("handler function called-" + name + "\n")
		h(writer,request)
	}
}

func check(h http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		t := request.URL.Query().Get("token")
		if t != "3des"{
			return
		}
		f := func(n string,a int) (in inner){
			in = inner{name:n,age:a,desc:fmt.Sprintf("my name is %s and i am %d years old",n,a)}
			return
		}

		combineFunc("prefix",f)
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
	/**
	最小惊讶原则,如果匹配的url以/结尾,那么当请求的url是/time/a/b/c的时候,如果没有其他url匹配,那么最终会匹配到/time/绑定的处理器上面
	但是如果制定的url是/time的话,那么/time/a/b/c就不会匹配到它,最终会匹配到/绑定的处理器上面
	 */
	http.HandleFunc("/time/", log(check(timeNow)))//链式的调用

	server.ListenAndServe()
}
