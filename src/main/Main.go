package main

import (
	"../chapter8"
	"os"
)

func main(){
	//chapter8.Tips()

	if os.Args[1] == "1" {
		chapter8.StartTimeServer()
	}else {
		chapter8.ConnectServer()
	}
}
