package main

import (
	"fmt"
	"time"
)

func main() {
	var data = make(chan string, 3)
	var notify = make(chan struct{}, 1)
	var monitor = make(chan struct{}, 2)

	go receiver_for_range(notify,monitor,data)

	go sender_for_range(notify,monitor,data)

	<- monitor
	<- monitor
}

func sender_for_range(n chan<- struct{}, m chan<- struct{}, d chan<- string) {
	for _, ch := range []string{"a", "b", "c", "d", "e"} {
		d <- ch
		fmt.Printf("send string:%s\n", ch)
		time.Sleep(100 * time.Millisecond)
		if ch == "c"{
			// send a signal
			n <- struct{}{}
		}
	}

	close(d)
	fmt.Printf("done.\n")
	m <- struct{}{}
}

func receiver_for_range(n <-chan struct{}, m chan<- struct{},d <- chan string){
	<- n
	for s := range d{
		fmt.Printf("receive string:%s\n",s)
	}
	m <- struct{}{}
}
