package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/base64"
	"time"
)

func setCookie(w http.ResponseWriter,r *http.Request){
	//这种方式这只的cookie没有设置过期时间,所以是会话cookie,当client重启了之后就消失
	c1 := http.Cookie{
		Name:"first_cookie",
		Value:"Go Web Programming",
		HttpOnly:true,
	}

	c2 := http.Cookie{
		Name:"second_cookie",
		Value:"Manning Publications Go",
		HttpOnly:true,
	}

	//w.Header().Set("Set-Cookie",c1.String())
	//w.Header().Set("Set-Cookie",c2.String())
	log.Printf("first_cookie:%s\n",c1.String())
	log.Printf("second_cookie:%s\n",c2.String())
	//或者使用http.SetCookie方法
	http.SetCookie(w,&c1)
	http.SetCookie(w,&c2)
}

func getCookie(w http.ResponseWriter,r *http.Request){
	h := r.Header["Cookie"]
	log.Printf("client cookie:%s\n",h)
	fmt.Fprintln(w,h)

	//也可以直接使用http.Request.Cookie(name string)的方法直接获取对应的cookie
	c,err := r.Cookie("first_cookie")
	if err != nil{
		fmt.Fprintf(w,"can not get cookied named:first_cookie\n")
	}else {
		fmt.Fprintf(w,"r.Cookie(\"first_cookie\") -> %s\n",c)
	}

	for k,v := range  r.Header{
		fmt.Fprintf(w,"%s:%s\n",k,v)
	}
}

func setMessage(w http.ResponseWriter,r *http.Request){
	msg := []byte("Hello world")
	c := http.Cookie{
		Name:"flash",
		Value:base64.URLEncoding.EncodeToString(msg),//将byte进行Base64之后以string的方式返回
	}
	http.SetCookie(w,&c)
}

func showMessage(w http.ResponseWriter,r *http.Request){
	c,err := r.Cookie("flash")
	if err != nil{
		if err == http.ErrNoCookie{
			fmt.Fprintln(w,"No Message found")
		}
		fmt.Fprintf(w,"empty cookie named:flash\n")
	}else {
		rc := http.Cookie{
			Name:"flash",
			//下面两行代码的本意就是让cookie失效,相当于命令浏览器删除这个cookie
			MaxAge:-1,
			Expires:time.Unix(1,0),//设置为一个已经过去的时间
		}
		http.SetCookie(w,&rc)
		val,_ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w,string(val))
	}
}

func main(){
	server := http.Server{
		Addr:"localhost:8080",
	}

	http.HandleFunc("/cookie",setCookie)
	http.HandleFunc("/showCookie",getCookie)
	http.HandleFunc("/setMessage",setMessage)
	http.HandleFunc("/getMessage",showMessage)
	server.ListenAndServe()
}
