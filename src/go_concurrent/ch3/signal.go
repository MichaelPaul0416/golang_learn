package main

import (
	"os"
	"syscall"
	"os/signal"
	"fmt"
)

func main(){
	signRecv := make(chan os.Signal,1)
	// 希望自行处理的信号
	sigs := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	// 如果第二个参数为空的话，那么就自行处理所有的信号,kill/stop除外
	// 操作系统向当前进程发送指定信号时发出通知，将当前进程指定的信号放入通道中
	// 这样该函数的调用方就可以从signal接收通道中按顺序获取操作系统发来的信号并进行相对应的处理
	signal.Notify(signRecv,sigs...)
	// 从通道中接受数据
	for sig := range signRecv{
		fmt.Printf("received a signal:%s\n",sig)
	}
}
