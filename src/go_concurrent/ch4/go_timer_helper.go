package main

import (
	"time"
	"fmt"
)

func main() {
	// 定义一个2s后到期的定时器
	//timerTask()

	// timeout
	//timeout()

	receiveWithTimeOut()

	// 定时器，到期之后，重新开启一个goroutine，执行第二个入参方法
	s := make(chan struct{})
	time.AfterFunc(1*time.Millisecond, func() {
		fmt.Printf("call task after timeout\n")
		s <- struct{}{}
	})

	<-s

}

func timerTask() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("current time:%v,\n", time.Now())
	expire := <-timer.C
	fmt.Printf("expire time:%v\n", expire)
	fmt.Printf("stop:%v\n", timer.Stop())
}

func timeout() {
	var dc = make(chan int, 1)
	var signal = make(chan struct{}, 1)

	go func() {
		select {
		// 一直阻塞，直到有数据返回
		case p := <-dc:
			fmt.Printf("receive data:%d\n", p)
			// case <- time.After(2 * time.Second):
			// 上面这行代码的效果和下面这行代码的效果是一样的
		case <-time.NewTimer(2 * time.Second).C:
			fmt.Printf("time out\n")
			signal <- struct{}{}
			// 如果打开的话，那就不会等待超时，而是直接返回
			//default:
			//	fmt.Printf("done.\n")
			//	signal <- struct{}{}
		}
	}()
	<-signal
}

func receiveWithTimeOut() {
	dc := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			dc <- i
		}

		close(dc)
	}()

	timeout := time.Millisecond * 50
	var timer *time.Timer
	for {
		// 设置每次的接受时间是500 ms
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}

		select {
		case i, ok := <-dc:
			if !ok {
				fmt.Printf("end.\n")
				return
			} else {
				fmt.Printf("receive data:%d\n", i)
			}
		case <-timer.C:
			fmt.Printf("current receive timeout\n")
			return
		}
	}
}
