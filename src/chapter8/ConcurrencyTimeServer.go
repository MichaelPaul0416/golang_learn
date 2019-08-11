package chapter8

import (
	"time"
	"fmt"
	"net"
	"io"
	"strings"
	"bufio"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func Tips() {
	go spinner(100 * time.Millisecond) //新起一个运行
	const length = 45
	r := fib(length)
	fmt.Printf("\r fib(%d) = %d\n", length, r)
}

const address = "localhost:8080"
//时钟服务器
func StartTimeServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("listen at port:%d falied-->%v\n", 8080, err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept a connection failed:%s\n", err)
			continue
		}
		go handleConnection(conn)//来一个链接就交给一个子协程处理，不阻塞当前main协程
	}
}

func handleConnection(con net.Conn) {
	defer closeConn(con)

	for {
		_, err := io.WriteString(con, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Printf("conn error:%v\n", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func closeConn(con net.Conn) {
	fmt.Printf("close connection:%v\n", con)
	con.Close()
}


func echo(c net.Conn,str string,delay time.Duration){
	fmt.Fprintf(c,"\t%s\n",strings.ToUpper(str))
	time.Sleep(delay * time.Millisecond)
	fmt.Fprintf(c,"\t%s\n",str)
	time.Sleep(delay * time.Millisecond)
	fmt.Fprintf(c,"\t%s\n",strings.ToLower(str))
}

func handleConnectionWithTimes(c net.Conn){
	input := bufio.NewScanner(c)
	for input.Scan(){
		s := input.Text()
		if s == "EOF"{
			fmt.Printf("close channel:%s\n",c)
			break
		}
		go echo(c,input.Text(),1000)//这里加上go关键字的含义是，对于客户端发送的每一个消息，都交给一个协程进行处理，而不是在一个连接中串行处理客户端请求
	}
	closeConn(c)
}

func TimeServerWithSeveralResponse(){
	listener,err := net.Listen("tcp",address)
	if err != nil{
		fmt.Printf("start server error:%v\n",err)
		return
	}

	for{
		c,err := listener.Accept()
		if err != nil{
			fmt.Printf("channel error:%v\n",err)
			continue
		}

		go handleConnectionWithTimes(c)
	}
}