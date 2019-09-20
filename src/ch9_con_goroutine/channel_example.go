package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//syncWithWaitGroup()
	syncWithGoroutine()
}

func syncWithGoroutine(){
	w1,w2 := make(chan bool),make(chan bool)
	go func() {
		for i:=0;i<10;i++{
			time.Sleep(1 * time.Microsecond)
			fmt.Printf("%d ",i)
		}

		w1 <- true
	}()

	go func() {
		for i:= 'A';i<'A'+10;i++{
			time.Sleep(1 * time.Microsecond)
			fmt.Printf("%c ",i)
		}

		w2 <- true
	}()

	// 等待从w1,w2通道接受信息
	<- w1
	<- w2

	fmt.Printf("\ndone")
}

func syncWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go func() {
		num := <-ch
		fmt.Printf("receive:%d\n", num)
		wg.Done()
	}()
	ch <- 10
	wg.Wait()
}
