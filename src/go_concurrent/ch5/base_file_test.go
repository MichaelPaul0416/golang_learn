package main

import (
	"testing"
	"fmt"
	"sync"
)

func TestNewDataFile(t *testing.T) {
	df, err := NewDataFile("info.txt", FIXED_LENGTH)
	if err != nil {
		t.Error(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(10)
	start := int('a')
	for i := start; i < start+10; i++ {
		go func(i int) {
			defer wg.Done()
			wsn,err := df.Write([]byte{byte(i)})
			if err != nil{
				t.Error(err)
			}else {
				fmt.Printf("write success:%v,and wsn:%d\n",i,wsn)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("done\n")
}
