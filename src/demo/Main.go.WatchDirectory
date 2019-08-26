package main

import (
	"fmt"
	"flag"
	"../chapter8"
)

func main(){
	flagParse()

	//chapter8.TotalFileSize()

	chapter8.TotalFileSizeByProcess()
	fmt.Println("------------------")
	chapter8.ConcurrencyDirSize()
}

func flagParse() {
	flag.Parse()
	args := flag.Args()
	//获取命令行的输入参数
	if len(args) == 0 {
		fmt.Printf("no arg input\n")
	} else {
		fmt.Printf("input args:%v\n", args)
	}
}
