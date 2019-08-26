package main

import (
	"net/http"
	"web/ch2/db"
)

//login
func login(w http.ResponseWriter, r *http.Request) {
	//根据layout.html的文件排版，这三个参数的含义分别是布局页面[导航页面+明细页面]
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

//logout
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		warning(err, "Failed to get cookie")
		session := db.Session{Uuid: cookie.Value}
		session.DeleteByUuid()
	}
	http.Redirect(w, r, "/login", 302)
}

//singup
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		danger(err, "can't parse form")
	}
	user := db.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	if _, err := user.Create(); err != nil {
		danger(err, "can not create user")
	}
	http.Redirect(w, r, "/login", 302)
}
