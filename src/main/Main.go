package main

import (
	"../chapter8"
)

func main(){

	//b := chapter8.SelfCycle()
	//fmt.Printf("return:%t\n",b)

	chapter8.Worker()

	chapter8.DownloadBigFile()
}
