package web

import (
	"net/http"
	"fmt"
)

func handler(writer http.ResponseWriter,r *http.Request){
	fmt.Fprintf(writer,"Hello world,%s!",r.URL.Path[1:])
}

func StartServer(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}
