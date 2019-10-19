package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		fmt.Println("received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("received:", elem, "[receiver]")
			}else {
				break
			}
		}

		fmt.Println("stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _,elem := range []string{"a","b","c","d"}{
			strChan <- elem
			fmt.Println("send:" ,elem,"[sender]")
			if elem == "c"{
				syncChan1 <- struct{}{}
				fmt.Println("send a sync signal. [sender]")
			}
		}

		fmt.Println("wait 2 seconds... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	// main goroutine 等待，直到返回
	<- syncChan2
	<- syncChan2
}
