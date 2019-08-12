package chapter9

import (
	"fmt"
	"math/rand"
	"bytes"
)

var saveChan = make(chan int)
var readChan = make(chan int)

func saveMoney(m int) {
	fmt.Printf("save money:%d\n", m)
	saveChan <- m
}

func readBalance() int {
	return <-readChan
}

func teller() {
	//将一个变量受限于单个goroutine内部，只有这个goroutine能访问并且修改，其他的goroutine都通过通道访问
	var balance = 100
	for {
		select {
		case s := <-saveChan:
			balance += s
		case readChan <- balance:
			fmt.Printf("return query balance:%d\n", balance)
		}
	}
}

func StartBankDemo(save, read int) {
	//只有一个goroutine能管理这个函数，对应的启动的goroutine就叫监控goroutine
	go teller()

	for i := 0; i < save; i++ {
		go saveMoney(rand.Intn(100))
	}

	for j := 0; j < read; j++ {
		func() {
			b := readBalance()
			fmt.Printf("read balance:%d\n", b)
		}()
	}

}

//将共享变量在通道之间依次传递，这也是保证这个变量是安全的一种方式
type cake struct {
	state string
	i int
}

func (c cake) String() string {
	var buf bytes.Buffer
	buf.WriteString("cake:[")
	buf.WriteString(c.state)
	buf.WriteString(",")
	fmt.Fprintf(&buf,"%d",c.i)
	buf.WriteString("]")
	return buf.String()
}

//cooked接受一个cake类型的指针
func baker(cooked chan<- *cake,i int) {
	c := cake{"cooked-烘焙",i}
	fmt.Printf("send pointer cake{%s} to next step...\n", c)
	cooked <- &c
}

/**
ice 是一个只能用来发送*cake的单向通道
bk 是一个只能用来接收*cake的单向通道
 */
func ice(ice chan<- *cake, bk <-chan *cake) {
	//bk通道是一个只接收*cake的通道，所以range的迭代对象cake已经是一个指针了，可以直接修改
	for cake := range bk {
		cake.state = "ice" + cake.state
		ice <- cake
	}
}

func TestCake(i int) {
	producer := make(chan *cake)
	middle := make(chan *cake)
	for m := 0; m < i; m++ {
		go func(p int) {
			baker(producer,p)
		}(m)//将m作为临时变量传入，形参中p接收m的值

		go func() {
			ice(middle,producer)
		}()
	}

	for ic := range middle{
		fmt.Printf("final cake:{%s}\n",ic)
	}

}
