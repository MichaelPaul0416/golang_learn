package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

/**
功能比ServerMux更加齐全的多路复用器
 */

func startServer() {
	mux := httprouter.New()
	//restful 形式
	mux.GET("/hello/:name",hello)

	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler:mux,
	}

	server.ListenAndServe()
}

//第三个参数其实就是restful格式的url中的变量形式的封装[在这里是:name]
func hello(response http.ResponseWriter,r *http.Request,p httprouter.Params){
	fmt.Fprintf(response,"hello,%s!\n",p.ByName("name"))
}

func main() {
	startServer()
}
