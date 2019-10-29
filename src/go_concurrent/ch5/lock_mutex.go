package main

import (
	"sync"
	"fmt"
	"time"
)

func main(){
	//lockRace()
	//unlockPanic()
	readWriteLock()
}

func unlockPanic(){
	var mutex sync.Mutex
	// 1.8 golang之前的版本，可以通过recover来恢复panic
	// 1.8之后的话，无法通过recover来恢复通过重复unlock一个已经解锁的mutex
	defer func() {
		fmt.Println("try to recover the panic.")
		if p := recover();p != nil{
			fmt.Printf("recovered the panic(%#v).\n",p)
		}
	}()

	fmt.Println("lock the lock")
	mutex.Lock()
	fmt.Println("lock is locked")
	fmt.Println("unlock the lock")
	mutex.Unlock()
	fmt.Println("the lock is unlocked")
	fmt.Println("unlock the lock")
	mutex.Unlock()
}
// 多个goroutine竞争一个锁
func lockRace() {
	var mutex sync.Mutex
	fmt.Println("lock the lock(main)")
	mutex.Lock()
	fmt.Println("main get the lock")
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("lock the lock.(g%d)\n", i)
			mutex.Lock()
			fmt.Printf("the lock is locked.(g%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("unlock the lock.(main)")
	mutex.Unlock()
	fmt.Println("the lock is unlocked.(main)")
	time.Sleep(time.Second * 1)
}

func readWriteLock(){
	var rwlock sync.RWMutex
	for i:=0;i<3;i++{
		go func(i int) {
			fmt.Printf("lock for reading...%d\n",i)
			rwlock.RLock()
			fmt.Printf("get read lock...%d\n",i)
			time.Sleep(2 * time.Second)
			fmt.Printf("unlock for reading...%d\n",i)
			rwlock.RUnlock()
			fmt.Printf("release read lock...%d\n",i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	rwlock.Lock()
	fmt.Printf("get write lock...\n")

}
