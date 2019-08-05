package main

import (
	"../chapter8"
	"os"
	"time"
	"fmt"
)

func main() {
	if len(os.Args) == 1 {
		//chapter8.Calculator(false)
		m := make(chan int)
		n := make(chan int)
		go chapter8.Counter(m, 3)
		go chapter8.ReceiveAndCalculator(n, m)
		chapter8.FinalResult(n)

		fmt.Printf("------------------------------------\n")
		chapter8.ProducerAndConsumer(3,6)

		time.Sleep(time.Hour * 1)

	} else {
		if os.Args[1] == "s" {
			chapter8.StartTimeServer()
		} else if os.Args[1] == "c" {
			chapter8.ClientToServer()
		}
	}

}
