package http

import (
	"net/http"
	"fmt"
)

func handler(writer http.ResponseWriter,r *http.Request){
	fmt.Fprintf(writer,"Hello world,%s!",r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}
