package lib

import (
	"testing"
	"fmt"
	"time"
	"sync"
)

func TestMyGoTickets_Take(t *testing.T) {
	gt, err := NewGoTickets(2)
	fmt.Printf("total:%d\n", gt.Total())
	if err != nil {
		fmt.Printf("init go tickets error:%v\n", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			defer func() {
				fmt.Printf("before return left:%d\n", gt.Remainder())
				gt.Return()
				fmt.Printf("after return left:%d\n", gt.Remainder())
				wg.Done()
			}()
			gt.Take()
			fmt.Printf("current cycle:%d\n", i)
			time.Sleep(1000 * time.Millisecond)
		}(i)
	}

	wg.Wait()

}
