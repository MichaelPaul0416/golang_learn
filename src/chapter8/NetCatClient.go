package chapter8

import (
	"net"
	"fmt"
	"io"
	"os"
)

//类似与telnet的客户端
func ConnectServer(){
	conn,err := net.Dial("tcp",address)
	if err != nil{
		fmt.Printf("connect server error:%s\n",err)
		return
	}

	defer closeConn(conn)

	if _,err := io.Copy(os.Stdout,conn);err != nil{
		fmt.Printf("write to stdout error:%s\n",err)
	}
}


func ConnectServerWithSend(){
	conn,err := net.Dial("tcp",address)
	if err != nil{
		fmt.Printf("connect server failed:%s\n",err)
		return
	}

	defer closeConn(conn)

	go func() {
		if _,err:=io.Copy(os.Stdout,conn);err != nil{
			fmt.Printf("receive and print error:%s\n",err)
		}
	}()

	if _,err := io.Copy(conn,os.Stdin);err != nil{
		fmt.Printf("read send message error:%s\n",err)
		return
	}
}

