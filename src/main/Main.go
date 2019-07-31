package main

import (
	"os"
	"../chapter8"
)

func main(){
	//chapter8.Tips()

	if os.Args[1] == "common"{
		if os.Args[2] == "1" {
			chapter8.StartTimeServer()
		}else {
			chapter8.ConnectServer()
		}
	}else if os.Args[1] == "special" {
		if os.Args[2] == "1"{
			chapter8.TimeServerWithSeveralResponse()
		}else {
			chapter8.ConnectServerWithSend()
		}
	}

}
