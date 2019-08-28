package main

import (
	"net/http"

	"fmt"
)

//go 1.6版本以上的话,https默认将使用http/2.0的版本

type Myhandler struct {}

func (h *Myhandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello world!")
}

func main(){
	handler := Myhandler{}

	server := http.Server{
		Addr:"127.0.0.1:8080",
		Handler:&handler,
	}

	fmt.Printf("server:%s\n",server)

}