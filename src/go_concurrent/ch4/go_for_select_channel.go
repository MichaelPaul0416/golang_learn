package main

import (
	"time"
	"fmt"
)

func main() {
	dc := make(chan int, 10)
	mc := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		dc <- i
		time.Sleep(100 * time.Millisecond)
	}

	close(dc)

	var nc chan struct{}
	fmt.Printf("nc == nil --> :%t\n",nc == nil)
	go func() {
	Loops:
		for {
			// 如果没有上面的Loops:和for循环的话，这里的select只会执行一次
			select {
			case num, ok := <-dc:
				{
					// 通道中没有数据
					if !ok {
						fmt.Printf("receive end\n")
						// break for 循环
						break Loops
					}
					fmt.Printf("receive data:%d\n", num)
				}
				case <-nc:{
					fmt.Printf("never call this line\n")
				}
			}
		}
		mc <- struct{}{}
	}()

	<-mc
}
