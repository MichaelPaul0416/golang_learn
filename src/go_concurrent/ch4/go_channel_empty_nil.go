package main

import (
	"fmt"
	"time"
)

// 下面的这个代码会报deadlock，因为当向一个没有初始化的nil chan发送数据的时候，会引起当前goroutine的永久阻塞
func main(){
	// 发送消息到nil通道
	//sendToNilChan()

	// send a message to a channel closed
	// cause a panic --> panic: send on closed channel
	sendToClosedChan()
}
func sendToClosedChan(){
	// signal
	signalChan := make(chan struct{},1)
	var channel = make(chan string,1)
	go func() {
		time.Sleep(100 * time.Millisecond)
		channel <- "hello"
		fmt.Printf("send done\n")
	}()

	close(channel)
	<- signalChan
}
func sendToNilChan() {
	// set a signal
	signalChan := make(chan struct{}, 1)
	var nilChan chan string
	go func() {
		// send a message to nil chan
		fmt.Printf("send a message to a nil chan\n")
		nilChan <- "hello"
		fmt.Printf("send done.")
		signalChan <- struct{}{}
	}()
	<-signalChan
}
