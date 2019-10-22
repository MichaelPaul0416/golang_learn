package main

import (
	"fmt"
	"time"
)

func main(){
	var mainSignal = make(chan struct{},2)
	// 等发送方发送完了，才开始接受的通道
	var receiveSignal = make(chan struct{},1)

	var dataChan = make(chan int,5)

	// receiver
	go func() {
		// wait until sender send a signal after close chan
		<- receiveSignal
		for{
			if data,ok := <- dataChan;ok{
				fmt.Printf("receive data:%d\n",data)
			}else {
				break
			}
		}
		// receive done
		fmt.Printf("receive all data\n")
		mainSignal <- struct{}{}
	}()

	// sender
	go func() {
		for i:=0;i<5;i++{
			dataChan <- i
			time.Sleep(100 * time.Millisecond)
		}

		close(dataChan)

		receiveSignal <- struct{}{}
		fmt.Printf("close data channel\n")
		mainSignal <- struct{}{}
	}()

	<- mainSignal
	<- mainSignal
}
