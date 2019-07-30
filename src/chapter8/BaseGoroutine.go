package chapter8

import (
	"time"
	"fmt"
)

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Fib(index int) int {
	if index < 2 {
		return index
	}

	return Fib(index-1) + Fib(index-2)
}
