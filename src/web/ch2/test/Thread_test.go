package test

import (
	"testing"
	"web/ch2/db"
	"fmt"
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
