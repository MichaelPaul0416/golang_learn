package main

import (
	"net"
	"strings"
	"fmt"
	"time"
	"io"
	"bytes"
	"math"
	"strconv"
	"math/rand"
	"sync"
)

/**
go 的socketAPI在操作系统底层的api上，使用的是非阻塞的api，也就是说go再调用操作系统底层的read,write方法的时候
是非阻塞的，但是go对我们开发者来说，将非阻塞的api最终封装成了阻塞的情况
 */
const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8080"
	DELIMITER      = '\t'
)
var wg sync.WaitGroup
func main() {

	wg.Add(1)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

func serverGo() {
	defer wg.Done()
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
		return
	}
	defer listener.Close()
	printServerLog("Got listener for the server. (local address: %s)", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept Error: %s", err)
		}
		printServerLog("Established a connection with a client application. (remote address: %s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// 先从连接中读取数据
	for {
		// 设置从当前连接中读取数据的超时时间
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("the connection is closed by another side.")
			} else {
				printServerLog("read error:%s", err)
			}
			break
		}

		printServerLog("received request:%s", strReq)

		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("send error message (written %d bytes):%s", n, err)
			continue
		}

		floatResp := math.Cbrt(float64(intReq))
		respMsg := fmt.Sprintf("the cube root of %d is %f.", intReq, floatResp)
		// 针对每一个请求，都设置10ms的时间处理，
		time.Sleep(10 * time.Millisecond)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("write error:%s", err)
		}
		printServerLog("send response (written %d bytes):%s", n, respMsg)
	}
	defer conn.Close()
}

func clientGo(id int){
	defer wg.Done()
	conn,err := net.DialTimeout(SERVER_NETWORK,SERVER_ADDRESS,10 * time.Second)
	if err != nil{
		printClientLog(id,"dial error:%s",err)
		return
	}

	defer conn.Close()
	printClientLog(id,"connected to server.(remote address:%s,local address:%s)",conn.RemoteAddr(),conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	// 发送和接受之前设置超时时间
	// 从开始发送第一条数据到接收到最后一条数据，总用时不超过5ms
	conn.SetDeadline(time.Now().Add(50 * time.Millisecond))

	// send
	for i :=0;i<requestNumber;i++{
		req := rand.Int31()
		n,err := write(conn,fmt.Sprintf("%d",req))
		if err != nil{
			printClientLog(id, "write error:%s", err)
			continue
		}
		printClientLog(id,"send request(written %d bytes):%d.",n,req)
	}

	// receive
	for j :=0;j<requestNumber;j++{
		strResp,err := read(conn)
		if err != nil{
			if err == io.EOF{
				printClientLog(id,"The connection is closed by another side.")
			}else{
				printClientLog(id,"read error:%s",err)
			}
			break
		}

		printClientLog(id,"receive response:%s.",strResp)
	}
}

func strToInt32(str string) (int32, error) {
	i32, err := strconv.ParseInt(str, 10, 32)
	return int32(i32), err
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func read(conn net.Conn) (string, error) {
	// 防止从连接值中读取多余的数据，从而对后序的读取操作造成影响
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		// 如果conn已经关闭，那么这里会返回一个io.EOF的error
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		// 缓冲区的长度是1.每次读取一个字节
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}
func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]:%s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args...)
}

func printClientLog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args...)
}

//func demo(){
/**
net.Listen方法是面向流的，第一个参数值可以是tcp,tcp4,tcp6,unix,unixpacket;也就是说必须是基于流的协议
至于udp,udp4,udp6,unixgram是基于udp的，并不是基于流的，所以不应该作为参数值
 */

// 返回第一个参数代表监听器
//listener, err := net.Listen("tcp", "127.0.0.1:8888")

//if err != nil {
//	fmt.Printf("listen error:%v\n", err)
//	return
//}

// 等待客户端连接请求，调用监听器的accept方法的时候，流程会被阻塞，直到某个客户端程序与当前程序建立tcp连接
//conn, err := listener.Accept()

//var dataBuffer bytes.Buffer
//b := make([]byte, 10)
//for {
// 设置io操作的超时时间，时间是一个绝对时间，仅仅针对当前连接conn之上的io操作
//conn.SetDeadline(time.Now().Add(2 * time.Second))
//n, err := conn.Read(b)
//if err != nil {
// 返回的异常类型如果是io.EOF的话，代表连接已经被远程关闭，在这个tcp连接之上已经没有可以再读取的数据，这个tcp连接可以被关闭
//if err == io.EOF {
//	fmt.Printf("connection closed by remote point\n")
//	conn.Close()
//} else {
//	fmt.Printf("read data error:%v\n", err)
//}
//break
//}
//dataBuffer.Write(b[:n])
//content := string(b[:n])
//fmt.Printf("receive:%s\n", content)
//}

// 返回协议：tcp,udp等
//conn.RemoteAddr().Network()
// 返回上述协议下的地址
//conn.RemoteAddr().String()
/**
net.Dial方法不向net.Listen方法一样，只能基于流的协议，它也可以基于数据包，所以第一个参数支持的传输层协议
除了net.Listen支持的协议之外，还支持udp,udp4,udp6,ip,ip4,ip6,unixgram等
客户端与server连接
 */
//net.Dial("tcp","127.0.0.1:8888")
// 带有超时的连接,默认单位是纳秒
//net.DialTimeout("tcp","127.0.0.1:8888",10 * time.Second)

//}
