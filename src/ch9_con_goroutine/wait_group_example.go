package main

import (
	"sync"
	"time"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	go printLetters(&wg)
	go printNumber(&wg)

	wg.Wait()
}

func printNumber(wg *sync.WaitGroup){
	for i:=0;i<10;i++{
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ",i)
	}

	wg.Done()
}

func printLetters(wg *sync.WaitGroup){
	for i:='A';i<'A' + 10;i++{
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ",i)
	}
	wg.Done()
}