package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int,1)
func main(){

	monitor := make(chan struct{},2)

	go func() {
		for{
			if elem,ok := <- mapChan;ok{
				elem["count"] ++
			}else {
				break
			}
		}
		fmt.Println("stopp. [receiver]")
		monitor <- struct{}{}
	}()

	go func() {
		countMap := make(map[string] int)
		for i:=0;i<5;i++{
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			// 引用类型作为通道传递的数据，接收方对引用的改变，会体现在发送方
			fmt.Printf("the count map:%v. [sender]\n",countMap)
		}
		close(mapChan)
		monitor <- struct{}{}
	}()

	<- monitor
	<- monitor
}
