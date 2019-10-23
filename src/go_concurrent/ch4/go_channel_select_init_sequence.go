package main

import "fmt"

var ch1 chan int
var ch2 chan int
var list = []chan int{ch1, ch2}
var nums = []int{1, 2, 3, 4, 5}

func main() {
	/**
	select的case表达式的执行语句的顺序是从左到右，从上到下
	同时由于这些chan都没有被初始化，所有向没有被初始化的chan发送数据，都会导致当前goroutine阻塞
	所以case1 case2都不会被执行，只会选择default执行，执行完毕之后直接返回
	 */
	select {
	case getChannel(0) <- getNumber(0):
		fmt.Printf("case 1\n")
	case getChannel(1) <- getNumber(1):
		fmt.Printf("case 2\n")
	default:
		fmt.Printf("default case\n")
	}

	chooseCase()

}

func getChannel(n int) chan int {
	fmt.Printf("choose channel:%d\n", n)
	return list[n]
}

func getNumber(i int) int {
	fmt.Printf("choose number index:%d\n", i)
	return nums[i]
}

func chooseCase() {
	ch := make(chan int, 5)
	// 当所有的case都可以执行的时候，从中随机选取一个case分支执行
	for i := 0; i < 5; i++ {
		select {
		case ch <- 1:
		case ch <- 2:
		case ch <- 3:
		}
	}

	for i:=0;i<5;i++{
		fmt.Printf("receive data:%d\n",<-ch)
	}
}
