package chapter8

import (
	"net"
	"io"
	"time"
	"fmt"
	"log"
)

//并发时钟服务器
func StartServer(){
	listener,err := net.Listen("tcp","localhost:8080")
	if err != nil{
		fmt.Printf("create listen err:%s\n",err)
		log.Fatal(err)
	}

	for{
		conn,err := listener.Accept()
		if err != nil{
			fmt.Printf("receive conn err:%s\n",err)
			log.Print(err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(c net.Conn){//类似socket
	defer c.Close()

	for{
		_,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil{
			fmt.Printf("err:%s\n",err)
			return
		}
		time.Sleep(time.Second * 1)
	}
}