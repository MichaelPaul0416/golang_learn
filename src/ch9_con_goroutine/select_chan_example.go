package main

import (
	"fmt"
	"time"
)

func senderA(ch chan string) {
	ch <- "message A"
}

func senderB(ch chan string) {
	ch <- "message B"
}

func main() {
	/**
	可以使用close关闭通道
	使用close函数关闭通道,相当于向通道的接收者表明该通道将不会再接受到任何值
	单向通道中只能执行接收操作的通道无法被关闭
	尝试向一个已经关闭的通道发送信息会引发一个panic,尝试关闭一个已经被关闭的通道也是如此
	尝试从一个已经被关闭的通道取值,总是会得到一个与通道类型想对应的零值,因此,从已经关闭的通道取值并不会导致goroutine被阻塞
	 */

	// 关闭通道的作用是通知其他正在尝试从这个通道接收值的goroutine,这个通道已经不会再接受到任何值
	//simpleSelect()

	channelClose()
}

func senderAWithClose(ch chan<- string) {
	ch <- "message A"
	close(ch)
}

func senderBWithClose(ch chan<- string) {
	ch <- "message B"
	close(ch)
}

func channelClose() {
	ch1, ch2 := make(chan string), make(chan string)

	go senderAWithClose(ch1)
	go senderBWithClose(ch2)

	ok1, ok2 := true, true
	var msg string
	for ok1 || ok2 {
		select {
		// 这里不能使用:=赋值,不然就相当于这两个变量是一次for循环中在当前case的临时变量,需要用全局变量
		// 当通道被关闭之后,这里的ok1的赋值就是false
		case msg, ok1 = <-ch1:
			if ok1 {
				fmt.Printf("receive from senderA:%s\n", msg)
			}
		case msg, ok2 = <-ch2:
			if ok2{
				fmt.Printf("receive from senderB:%s\n", msg)
			}
		default:
			fmt.Printf("default branch\n")
		}
	}

}

func simpleSelect() {
	cha, chb := make(chan string), make(chan string)
	/**
	下面两行注释掉之后，会发生死锁
	但是如果加上了default分支,就不会发生死锁
	 */
	go senderA(cha)
	go senderB(chb)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case m := <-cha:
			fmt.Printf("receive from A:%s\n", m)
		case m := <-chb:
			fmt.Printf("receive from B:%s\n", m)
			// 一般来说select里面都要添加default分支,这样不容易发生deadlock
		default:
			fmt.Printf("default branch\n")
		}

	}
}
