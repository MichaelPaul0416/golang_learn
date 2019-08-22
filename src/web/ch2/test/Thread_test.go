package test

import (
	"testing"
	"web/ch2/db"
	"fmt"
	"time"
	"log"
)

func TestThread(t *testing.T){

	ts,err := db.Threads()
	if err != nil{
		fmt.Printf("error:%v\n",err)
		return
	}

	for _,t := range ts{
		fmt.Printf("%s\n",t)
	}
}


func TestDatetimeNow(t *testing.T){
	fmt.Printf(db.DateTimeNow())
}

func TestInsertSession(t *testing.T){
	uuid := db.CreateUUID()
	user := db.User{Id: 2, Uuid: uuid, Name: "Paul", Email: "sdbfduwv@88.com", Password: "123456", CreatedAt: time.Now()}
	session,err := user.CreateSession()
	if err != nil{
		log.Printf("create session error:%v\n",err)
		return
	}

	fmt.Printf("session:%s\n",&session)
}