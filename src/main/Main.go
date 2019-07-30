package main

import (
	"../chapter8"
)

func main() {
	//go chapter8.Spinner(100 * time.Millisecond)//新开一个协程
	//const num = 45
	//res := chapter8.Fib(num)
	//fmt.Printf("Fib result:%d\n",res)

	//
	chapter8.StartServer()
}
