package chapter8

import (
	"fmt"
	"time"
	"os"
)

func Rocket() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 3; countdown > 0; countdown-- {
		fmt.Println(countdown)
		//从tick通道中接收消息
		t := <-tick
		fmt.Printf("receive from empty channel:%s\n", t)
	}

	fmt.Println("rocket launch...")
}

func RocketWithSelect() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		//从输入中读取一个个byte之后,马上发送终止的消息给abort通道
		abort <- struct{}{}
	}()

	fmt.Println("send message after a time")
	//select接收的其实是一次通讯
	select {
	case <-time.After(1 * time.Second): //马上返回一个通道,并且启动一个新的goroutine,然后3s之后发送消息到这个通道上,让select接收
		fmt.Println("time is up...")
	case <-abort:
		fmt.Printf("cancel...\n")
		return
	}
	fmt.Printf("send message now...\n")
}

//利用带缓冲区的通道,输出偶数
func ShowNumber(l int, b bool) {
	ch := make(chan int, 1) //缓冲区长度为1的chan
	//由于ch此时是空的,所以从它这里无法接收,即便是带有缓冲区的,如果打开下面这个注释,那么将会永久阻塞
	//fmt.Printf("p:%d\n",<- ch)
	for i := 0; i < l; i++ {
		select {
		case ch <- i:
			fmt.Printf("send:%d\n", i)
		case x := <-ch:
			fmt.Printf("num:%d\n", x)
		}
	}
	if b {
		close(ch)
	}

	//当通道中的所有数据都发送并且接受完毕了,再接收的话
	// -- 如果通道已经关闭,返回通道对应类型个的零值
	// -- 如果通道还没有关闭,那么永久阻塞
	fmt.Printf("recieve after all done:%d\n", <-ch)
}

//当select中有多个case可以被执行时,随机选择一个case执行
func MultiCasesSelector(i int) {
	ch := make(chan int, 3)
	send, receive := 0, 0
	for n := 0; n < i; n++ {
		select { //并不是一次for循环执行一次select的匹配,而是开启一个select,也就是说最后会有10个select,而不是一次for循环中去匹配哪个case可以执行
		case ch <- n:
			fmt.Printf("[%d]send:%d\n", n, n)
			send++
		case x := <-ch:
			fmt.Printf("[%d]receive:%d\n", n, x)
			receive++
		}
	}
	//此时send>=receive的,因为如果缓冲区空了[此时send=receive],那么此次能执行的情况一定只有send,也就是说send肯定是大于等于receive的
	fmt.Printf("select done:send[%d]/receive[%d]\n", send, receive)
	if send > receive { //send的次数大于receive
		for left := range ch {
			fmt.Printf("left:%d\t", left)
			receive ++
			if receive == send {
				close(ch)
				//break
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("all done:send[%d]/receive[%d]\n", send, receive)
}

func SendMessageWithTick(t int) {

	/**
	* 使用tick的时候,推荐使用如下的模式:好处是相比time.Tick生成的tick对象,下面这种方式可以通过Stop手动关闭,避免在函数执行完毕,返回上层函数之后,time.Tick生成的tick还每秒都接受数据
	* tick := time.NewTicker(1 * time.Second)
	* <- tick.C
	* tick.Stop()//手动关闭
	*
	 */

	//开启每秒都会写入一个time的通道
	//tick := time.Tick(1 * time.Second)
	tick := time.NewTicker(1 * time.Second) //这个对象有一个Stop方法,可以手动关闭tick
	abort := make(chan struct{})

	//这个goroutine要写在for前面,不然的执行for的是main goroutine会等for执行完了才去执行go代码
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for i := t; i > 0; i-- {
		select {
		case <-tick.C:
			fmt.Printf("send message left time:%d\n", i)
		case <-abort:
			fmt.Printf("cancel send message\n")
			return
		}
	}

	//当这里返回的时候
	fmt.Printf("send message:hello\n")

	//虽然上面的select已经返回了,但是tick通道还没关闭,还是秒都会发送消息到tick中,所以下面的代码还是可以继续接收
	i := 0
	for p := range tick.C {
		fmt.Printf("tick still alive:%v\n", p)
		i++
		if i > 3 {
			tick.Stop()//手动关闭,可以防止tick的内存泄漏,通道开着每秒都一直输出
			break
		}
	}
}

func UnBlockingSelector(){
	abort := make(chan  struct{})

	go func() {
		os.Stdin.Read(make([]byte,1))
		abort <- struct{}{}
	}()

	//1.如果没有for死循环的话,在没有abort接收数据的前提下,default只会执行一遍,就马上结束
	//2.如果没有for死循环,如果没有default的话,那么代码会一直阻塞在select上
	//3.如果有for死循环,但是没有default的话,那么代码会阻塞在for的第一次循环时,卡在第一次循环的select上
	//4.如果有for死循环,并且有default,同时abort中没有数据读进来的话,那么default中的代码快就会每秒都执行一次
	for {
		select {
		case <- abort:
			fmt.Printf("cancel...\n")
			return
		default:
			fmt.Printf("un blocking...\n")
			time.Sleep(time.Second * 1)
		}
	}
}