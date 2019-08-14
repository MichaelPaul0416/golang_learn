package chapter9

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

/**
互斥锁的简单实现和golang提供的互斥锁的应用
 */


 /**
 之前提到了可以用互斥量,或者监视通道的方式来实现并发安全,现在也可以使用封装来实现并发安全
 具体如下
  */
var l sync.Mutex

type Account struct{
	name string
	bala int
}

func (a Account)  addBalance(b int){
	a.bala += b
}

func (a Account) getBalance() int{
	return a.bala
}

/**
使用不导出的方法和结构体属性来保证了并发的安全性和事务
 */
func (a Account) AddBala(b int) bool{
	l.Lock()
	defer l.Unlock()

	//先扣款
	a.addBalance(-b)
	if a.getBalance() < 0{
		a.addBalance(b)
		fmt.Printf("has not enough money\n")
		return false
	}

	return true
}

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func putMoney(b int, w *sync.WaitGroup) {
	defer w.Done()
	sema <- struct{}{}
	fmt.Printf("put money:%d/total:%d\n", b, b+balance)
	balance += b
	<-sema
}

func getMoney(w *sync.WaitGroup) int {
	defer w.Done()
	sema <- struct{}{}
	b := balance
	fmt.Printf("get money:%d\n", balance)
	<-sema
	return b
}

func CustomExclusiveLock(p, g int) {
	var watch sync.WaitGroup
	for i := 0; i < p; i++ {
		b := rand.Intn(10)
		watch.Add(1)
		go putMoney(b, &watch)
	}

	for i := 0; i < g; i++ {
		watch.Add(1)
		go getMoney(&watch)
	}

	go func() {
		watch.Wait()
		close(sema)
		fmt.Printf("done and final balance:%d...\n", balance)
	}()

	for {
		time.Sleep(10 * time.Second)
		break
	}
}

/**
使用golang自带的互斥锁实现
 */

var lock sync.Mutex
func putMoneyWithLock(b int,w *sync.WaitGroup){
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf("put money:%d/total:%d\n",b,b+ balance)
	balance += b
	defer w.Done()
}

func getMonetWithLock(w *sync.WaitGroup) int{
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf("get money:%d\n",balance)
	defer w.Done()
	return balance
}

func ConcurrencyWithLock(p,g int){
	var w sync.WaitGroup
	for i:=0;i<p;i++{
		w.Add(1)
		m := rand.Intn(10)
		go putMoneyWithLock(m,&w)
	}

	for i:=0;i<g;i++{
		w.Add(1)
		go getMonetWithLock(&w)
	}

	go func() {
		w.Wait()
		fmt.Printf("finally money:%d\n",balance)
	}()

	for {
		time.Sleep(10 * time.Second)
		break
	}
}


/**
读写锁的使用demo
 */
type Num int

var number Num

var rw sync.RWMutex
func writeNumber(p Num,w *sync.WaitGroup){
	rw.Lock()
	defer rw.Unlock()
	defer w.Done()
	fmt.Printf("write:%d\n",p)
	number = p
}

func readNumber(w *sync.WaitGroup) Num{
	rw.RLock()
	defer rw.RUnlock()
	defer w.Done()
	return number
}

func ReadWriteLock(r,w int) {
	var watch sync.WaitGroup
	for i:=0;i<r;i++{
		watch.Add(1)
		go func() {
			fmt.Printf("read:%d\n",readNumber(&watch))
		}()
	}

	for i:=0;i<w;i++{
		watch.Add(1)
		go func(m int) {
			writeNumber(Num(m),&watch)
		}(i)
	}

	go func() {
		watch.Wait()
		fmt.Printf("read write lock done and Num:%d...\n",number)
	}()

	for {
		time.Sleep(10 * time.Second)
		break
	}
}
