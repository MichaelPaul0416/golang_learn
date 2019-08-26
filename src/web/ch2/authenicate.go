package main

import (
	"net/http"
	"web/ch2/db"
	"log"
)

/**
校验
 */

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := db.UserByEmail(r.PostFormValue("email")) //获取表单中提交的email字段的值
	if err != nil{
		danger(err,"can not find by email")
		return
	}

	if user.Password == db.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Printf("create user session error:%v\n", err)
			w.Write([]byte("create user session error"))
			return
		}

		cookie := http.Cookie{
			Name:   "_cookie",
			Value: session.Uuid,
			HttpOnly:true,
		}

		http.SetCookie(w,&cookie)
		http.Redirect(w,r,"/",302)
		return
	}

	http.Redirect(w,r,"/login",302)
}
