package chapter1

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

/**
 * 简单的web服务器
 */

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/http_info", http_info)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func http_info(response http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(response,"%s %s %s\n",r.Method,r.URL,r.Proto)
	for k,v := range r.Header{
		fmt.Fprintf(response,"Header[%q] = %q\n",k,v)
	}

	fmt.Fprintf(response,"Host = %q\n",r.Host)
	fmt.Fprintf(response,"RemoteAddr = %q\n",r.RemoteAddr)
	if err := r.ParseForm(); err != nil{
		log.Print(err)
	}

	for k,v :=range r.Form{
		fmt.Fprintf(response,"Form[%q] = %q\n",k,v)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

func counter(response http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(response, "count : %d\n", count)
	mu.Unlock()
}
