package test

import (
	"testing"
	"web/ch2/db"
	"fmt"
	"time"
	"log"
)

var user = db.User{
	Id: 2,
}

var sessionExample = db.Session{
	Uuid:"8ad931d8-a9d1-4477-5357-38e80ecc5b1b",
}

func TestThread(t *testing.T) {

	ts, err := db.Threads()
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}

	for _, t := range ts {
		fmt.Printf("%s\n", t)
	}
}

func TestDatetimeNow(t *testing.T) {
	fmt.Printf(db.DateTimeNow())
}

func TestInsertSession(t *testing.T) {
	uuid := db.CreateUUID()
	user := db.User{Id: 2, Uuid: uuid, Name: "Paul", Email: "sdbfduwv@88.com", Password: "123456", CreatedAt: time.Now()}
	session, err := user.CreateSession()
	if err != nil {
		log.Printf("create session error:%v\n", err)
		return
	}

	fmt.Printf("session:%s\n", &session)
}

func TestSessionQuery(t *testing.T) {
	s,e := user.Session()
	if e != nil{
		log.Fatal(e)
		return
	}
	fmt.Printf("session:%s\n",&s)
}


func TestCheckValid(t *testing.T){
	b,e := sessionExample.CheckValid()
	if e != nil{
		log.Fatal(e)
		return
	}
	fmt.Printf("session valid:%t\n",b)
}

func TestDeleteByUuid(t *testing.T){
	if err := sessionExample.DeleteByUuid();err != nil{
		log.Fatal(err)
	}
	fmt.Printf("delete ok")
}

func TestCreateUser(t *testing.T){
	user := db.User{Name:"Bob",Email:"45789@qq.com",Password:"dfho"}

	if _,err := user.Create();err != nil{
		log.Printf("create user error:%v\n",err)
		return
	}

	log.Printf("new user id:%d\n",user.Id)
}

func TestDeleteUser(t *testing.T){
	user := db.User{Id:1}
	if err := user.Delete();err != nil{
		log.Printf("delete user failed:%v\n",err)
		return
	}

	log.Printf("delete ok")
}

func TestFindAllUsers(t *testing.T){
	var ary []db.User
	var err error
	if ary,err = db.Users();err != nil{
		log.Printf("query users failed:%s\n",err)
		return
	}

	for _,user := range ary{
		log.Printf("user:%s\n",&user)
	}
}

func TestCreateThread(t *testing.T){
	th,err := user.CreateThread("coding")
	if err != nil{
		t.Errorf("create thread failed:%v\n",err)
		return
	}
	log.Printf("create thread successfully:%s\n",th)

	p,err := user.CreatePost(th,"good coding")
	if err != nil{
		t.Errorf("create post for thread[uuid:%s]:%v\n",th.Uuid,err)
		return
	}
	log.Printf("create post successfully:%s\n",&p)
}