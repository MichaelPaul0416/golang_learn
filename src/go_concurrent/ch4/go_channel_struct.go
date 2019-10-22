package main

import (
	"time"
	"fmt"
)

type wrapper struct {
	count int
}

type changeWrapper struct {
	count int
}

var structChan = make(chan map[string]wrapper, 1)
var changeChan = make(chan map[string]*changeWrapper,1)

func main() {
	signalChan := make(chan struct{}, 2)
	signalChangeChan := make(chan struct{},2)
	// receive
	go func() {
		for{
			if elem,ok := <- structChan;ok{
				counter := elem["count"]
				counter.count++
			}else {
				break
			}
		}

		fmt.Printf("stop [receiver.]\n")
		signalChan <- struct{}{}
	}()

	go func(){
		for{
			if ch,ok := <-changeChan ; ok{
				cw := ch["count"]
				cw.count ++
			}else {
				break
			}
		}
		fmt.Printf("pointer stop [pointer receiver.]\n")
		signalChangeChan <- struct{}{}
	}()

	// send
	go func() {
		// 构建一个空的map
		countMap := map[string]wrapper{
			"count": wrapper{},
		}

		for i := 0; i < 5; i++ {
			structChan <- countMap

			time.Sleep(100 * time.Millisecond)
			fmt.Printf("the count map:%v, [sender]\n",countMap)
		}
		close(structChan)
		signalChan <- struct{}{}
	}()

	go func() {
		// 即便通道传输的是map，引用类型，但是map的value如果是值类型的话，接收方对其的改动不会影响到发送方
		// 只有map的value是引用类型（指针）的时候，接收方对其的改动才会影响到发送方
		chMap := map[string]*changeWrapper{
			"count": &changeWrapper{},
		}
		for i:=0;i<5;i++{
			changeChan <- chMap
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("pointer the count map:%d, [pointer sender]\n",chMap["count"].count)
		}

		close(changeChan)
		signalChangeChan <- struct{}{}
	}()

	<-signalChan
	<-signalChan
}
