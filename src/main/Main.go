package main

import (
	"../chapter9"
	"sync"
)

func main() {
	chapter9.Memory_1()

	//multiInitByRwlock()

	multiInitBySyncOnce()

}

func multiInitBySyncOnce() {
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(m int) {
			defer w.Done()
			s := id(m)
			chapter9.InitAndQueryBySyncOnce(s)
		}(i)
	}
	w.Wait()
}

func multiInitByRwlock() {
	var wait sync.WaitGroup
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(m int) {
			defer wait.Done()
			s := id(m)
			chapter9.PersonInfo(s)
		}(i)
	}
	wait.Wait()
}

func id(m int) string {
	var s string
	if m%2 == 0 {
		s = "001"
	} else {
		s = "002"
	}
	return s
}
