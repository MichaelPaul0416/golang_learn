package main

import (
	"net/http"
	"web/ch2/db"
	"fmt"
)

//包含thread的一些操作，业务层面
func newThread(w http.ResponseWriter, request *http.Request) {
	_, err := session(w, request)
	if err != nil {
		http.Redirect(w, request, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}

	err = r.ParseForm()
	if err != nil {
		danger(err, "can't parse form")
	}

	user, err := s.User()
	if err != nil {
		danger(err, "can not get user by session")
	}
	topic := r.PostFormValue("topic")
	var addThread db.Thread
	if addThread, err = user.CreateThread(topic); err != nil {
		danger(err, "can not create a thread")
	} else {
		info("create a thread:%s\n", &addThread)
	}
	http.Redirect(w, r, "/", 302)
}

func readThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := db.ThreadByUuid(uuid)
	if err != nil {
		error_message(w, r, "can not read thread by uuid:"+uuid)
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

//add a new post for thread
func postThread(writer http.ResponseWriter, r *http.Request) {
	s, err := session(writer, r)
	if err != nil {
		http.Redirect(writer, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger(err, "can not parse form")
		}
		user, err := s.User()
		if err != nil {
			danger(err, "can not get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := db.ThreadByUuid(uuid)
		if err != nil {
			error_message(writer, r, "can not read thread")
		}

		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "can not create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, r, url, 302)

	}
}
