package main

import (
	"fmt"
)

func senderA(ch chan string){
	ch <- "message A"
}

func senderB(ch chan string){
	ch <- "message B"
}

func main(){
	cha,chb := make(chan string),make(chan string)
	/**
	下面两行注释掉之后，会发生死锁
	 */
	//go senderA(cha)
	//go senderB(chb)

	//for i:=0;i<5;i++{
		select {
		case m := <- cha:
			fmt.Printf("receive from A:%s\n",m)
		case m := <- chb:
			fmt.Printf("receive from B:%s\n",m)
		}

	//}
}
