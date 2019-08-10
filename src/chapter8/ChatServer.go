package chapter8

import (
	"fmt"
	"net"
	"log"
	"bufio"
)

//只能用于发送消息的单向通道
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func StartChatServer() {
	listener,err := net.Listen("tcp","localhost:8080")
	if err != nil{
		log.Fatal(err)
	}

	go broadcaster()

	for{
		conn,err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		fmt.Printf("receive a connection\n")
		go handleClientConnection(conn)
	}
}

//广播者,作为中间角色,负责监听entering/leaving/message三个通道的消息
func broadcaster() {
	clients := make(map[client]bool)
	//只能用来发现的chan作为key
	for {
		select {
		case msg := <-message:
			fmt.Printf("public message event\n")
			for cli := range clients {
				//公共消息发送给每个与client的通道,然后每个与client的链接会在clientWriter中接受chan中的消息,然后写回到connection中
				cli <- msg
			}
		case cli := <-entering:
			fmt.Printf("store a connection\n")
			clients[cli] = true
		case cli := <-leaving:
			fmt.Printf("remove a connection\n")
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleClientConnection(conn net.Conn) {
	//为与client的链接创建一个chan,用于该链接异步通过其异步发送消息
	ch := make(chan string)
	go clientWriter(conn, ch)

	from := conn.RemoteAddr().String()
	//将消息发送到与client的通道中
	ch <- "welcome remote client " + from + "\n"
	//通过message告诉监听者,有人连接进来了
	message <- fmt.Sprintf("remote client[%s] connected to server\n", from)

	//通过entering告诉监听者连入的client
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- from + ":" + input.Text()
	}

	//message <- fmt.Sprintf("%s : random int --> %d\n",from,rand.Intn(100))

	leaving <- ch
	message <- fmt.Sprintf("remote client[%s] leaving\n", from)
	conn.Close()
}

//每一个与client链接的chan异步发送
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}

//单向通道,只能用于接收消息,不能用于发送消息
func differ(ch <-chan int) {
	for msg := range ch {
		fmt.Printf("string:%d\n", msg)
	}
}
