package main

import (
	"../chapter7"
	"fmt"
)

func main(){
	var c chapter7.ByteCounter
	c.Write([]byte("hello"))
	fmt.Printf("imple:%v\n",c.String())

	//ByteCounter实际类型本身就是int
	c = 0
	var name = "Paul"
	//ByteCounter不需要声明实现了什么接口，只需要实现对应的Write方法，就说明它是Writer接口的实现
	fmt.Fprintf(&c,"hello,%s\n",name)
	fmt.Printf("impl:%v\n",c.String())

	var crw chapter7.CloseableReaderWriter
	crw.Write([]byte("write"))
	//crw.Read([]byte("read"))
	if err:=crw.Close();err == nil{
		fmt.Printf("crw closeable:%t\n",true)
	}else {
		fmt.Printf("crw closeable:%t\terror:%v\n",false,err)
	}
}