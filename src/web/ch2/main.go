package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"web/ch2/db"
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

	mux.HandleFunc("/",index)


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
	threads,err := db.Threads()
	if err != nil{
		fmt.Fprintf(response,"query all threads failed")
		return
	}
	log.Printf("total threads[%d]\n",len(threads))
	templates.ExecuteTemplate(response,"layout",threads)

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
