package ch2

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
)

func SimpleHttpServer() {
	//新建一个多路复用器
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("D:/GoLearn"))
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
	//slice
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	
}

func showHttpHeaderParam(response http.ResponseWriter,request *http.Request){
	for h := range request.Header{
		log.Printf("key:%s -- value:%s\n",h,request.Header.Get(h))
		fmt.Fprintf(response,"key:%s -- value:%s\n",h,request.Header.Get(h))
	}
	log.Printf("------------------------------------------\n")
}
