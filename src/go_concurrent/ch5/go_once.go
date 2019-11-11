package main

import (
	"sync"
	"fmt"
)

func main() {
	var once sync.Once
	for i := 0; i < 5; i++ {
		// 结果是只会输出一行，current:0
		once.Do(func() {
			fmt.Printf("current:%d\n",i)
		})
	}
}
