package main

import (
	"time"
	"fmt"
)

func main() {
	// set interval
	shortInterval := time.Second
	longInterval := time.Second * 2

	// set no buffer chan
	ch := make(chan int, 0)
	go func() {
		var t0, t1 int64
		for i := 0; i < 10; i++ {
			ch <- i
			t1 = time.Now().Unix()
			if t0 == 0 {
				fmt.Printf("first send:%d\n", i)
			} else {
				fmt.Printf("send number: %d after %d's\n", i, t1-t0)
			}
			t0 = time.Now().Unix()
			time.Sleep(shortInterval)
		}

		close(ch)
	}()

	var s0, s1 int64
Loop:
	for {
		select {
		case i, ok := <-ch:
			{
				if !ok {
					fmt.Printf("receive end!\n")
					break Loop
				}
				s1 = time.Now().Unix()
				if s0 == 0 {
					fmt.Printf("receive : %d first time\n", i)
				} else {
					fmt.Printf("receive: %d with a interval: %d\n", i, s1-s0)
				}
				s0 = time.Now().Unix()
				time.Sleep(longInterval)
			}
		}
	}

	fmt.Printf("task done\n")
}

