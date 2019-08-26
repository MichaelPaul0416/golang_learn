package main

import (
	"net/http"
	"fmt"
	"log"
	"web/ch2/db"
)

func SimpleHttpServer() {
	//新建一个多路复用器
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("./public/"))

	mux.HandleFunc("/",index)
	// set error page
	mux.HandleFunc("/err",err)

	mux.HandleFunc("/login",login)
	mux.HandleFunc("/logout",logout)
	mux.HandleFunc("/signup",signup)
	mux.HandleFunc("/signup_account",signupAccount)
	mux.HandleFunc("/authenticate",authenticate)


	mux.HandleFunc("/thread/new",newThread)
	mux.HandleFunc("/thread/create",createThread)
	mux.HandleFunc("/thread/post",postThread)
	mux.HandleFunc("/thread/read",readThread)

	//处理器
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//HandleFunc是将请求交给处理器函数
	mux.HandleFunc("/tips", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome:%s and visit path:%s\n", request.URL.Query().Get("name"), request.URL.Path)
		showHttpHeaderParam(writer,request)
	})



	server := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(response http.ResponseWriter,r *http.Request){
	if "favicon.ico" == r.URL.Path{
		return
	}

	_,err := session(response,r)

	threads,err := db.Threads()
	if err != nil{
		fmt.Fprintf(response,"query all threads failed")
		return
	}

	if err != nil{
		generateHTML(response,threads,"layout","public.navbar","index")
	}else{
		generateHTML(response,threads,"layout","private.navbar","index")
	}

}

func err(w http.ResponseWriter,r *http.Request){
	//获取get请求中的键值对
	vals := r.URL.Query()

	_,err := session(w,r)
	if err != nil{
		generateHTML(w,vals.Get("msg"),"layout","public.navbar","error")
	}else{
		generateHTML(w,vals.Get("msg"),"layout","private.navbar","error")
	}
}

func showHttpHeaderParam(response http.ResponseWriter,request *http.Request){
	for h := range request.Header{
		log.Printf("key:%s -- value:%s\n",h,request.Header.Get(h))
		fmt.Fprintf(response,"key:%s -- value:%s\n",h,request.Header.Get(h))
	}
	log.Printf("------------------------------------------\n")
}

func main(){
	SimpleHttpServer()
}
