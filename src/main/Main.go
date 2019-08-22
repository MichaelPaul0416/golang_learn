package main

import (
	"web/ch2/db"
	"fmt"
)

func main()  {
	ts,err := db.Threads()
	if err != nil{
		fmt.Printf("error:%v\n",err)
		return
	}

	for t := range ts{
		fmt.Printf("%v\n",t)
	}
}
