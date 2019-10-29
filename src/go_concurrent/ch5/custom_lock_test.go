package main

import (
	"testing"
	"fmt"
	"sync"
)

func TestCloseableChannel(T *testing.T) {
	ch := make(chan struct{})

	defer func() {
		var err error
		var ok bool
		if p := recover(); p != nil {
			err, ok = interface{}(p).(error)
			if ok {
				fmt.Printf("error:%v\n", err)
			} else {
				fmt.Printf("panic:%v\n", p)
			}
		}
	}()

	close(ch)
	ch <- struct{}{}
}

func TestClose(T *testing.T) {
	ch := make(chan struct{})
	close(ch)
	_,ok := TryCloseableChannel(ch)
	fmt.Printf("%t\n", ok)
}

func TestUnReentrantLock_Lock(t *testing.T) {
	lock, err := NewUnReentrantLock()
	if err != nil {
		fmt.Printf("generate lock error\n")
		return
	}

	i := 0

	var wg sync.WaitGroup
	wg.Add(20)
	for m := 0; m < 20; m++ {
		go func(i *int) {
			lock.Lock()
			defer func() {
				wg.Done()
				lock.UnLock()
			}()
			fmt.Printf("get lock:%d\n",*i)
			*i++
		}(&i)
	}

	wg.Wait()
	fmt.Printf("final :%d\n",i)
}

func TestBlockingChannel(t *testing.T){
	// 这里一定要长度为1，如果是没有长度的话，就是阻塞的通道，也就是执行发送操作之后，当前goroutine就会一直阻塞
	ch := make(chan struct{},1)

	ch <- struct{}{}

	fmt.Printf("send ok\n")

	// must be deadlock
	ch <- struct{}{}
	fmt.Printf("sencond send ok \n")

	<- ch
	fmt.Printf("done.\n")

}
