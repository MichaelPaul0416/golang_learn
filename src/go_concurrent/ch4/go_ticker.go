package main

import (
	"time"
	"fmt"
)

/**
声明一个断续器，每秒执行一次，通过timer.Ticker.C这个通道来接受，每次到期的通知
 */
func main() {
	tc := time.NewTicker(1 * time.Second)
	var i int
	// for + range:这种形式是从一个可缓冲的通道中，持续不断的接受值，直到接受完毕或者通道被关闭
Loop:
	for {
		select {
		case <-tc.C:
			fmt.Printf("time:%v\n", time.Now())
		// 以内部时间为准
			time.Sleep(2 * time.Second)
			if i < 10 {
				i++
			} else {
				break Loop
			}
		}
	}

	fmt.Printf("done.\n")
}
