package main

import (
	"fmt"
	"time"
)

// 单向通道的操作
func main() {
	var dataChan = make(chan int, 5)
	var mainSignal = make(chan struct{}, 2)
	var notifySignal = make(chan struct{}, 1)

	go receiver(notifySignal,dataChan,mainSignal)

	go sender(notifySignal,mainSignal,dataChan)

	<- mainSignal
	<- mainSignal
}

// 第一个和第二个参数是接收数据的单向通道，第三个参数是发送数据的单向通道
func receiver(notify <-chan struct{}, dc <-chan int, ms chan<- struct{}) {
	<-notify
	for {
		if d, ok := <-dc; ok {
			fmt.Printf("receive data:%d\n", d)
		} else {
			break
		}
	}
	fmt.Printf("receive done.\n")
	ms <- struct{}{}
}

func sender(notify chan<- struct{}, ms chan<- struct{}, dc chan<- int) {
	for i := 0; i < 5; i++ {
		dc <- i
		time.Sleep(100 * time.Millisecond)
	}
	notify <- struct{}{}
	close(dc)
	fmt.Printf("all send done.\n")
	ms <- struct{}{}
}
