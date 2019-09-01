package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

/**
文件的读写
 */

func main() {
 	data := []byte("Hello Golang\n")
 	//使用ioutil.WriteFile/ReadFile方法的时候,write的入参和read的出参都是slice
 	err := ioutil.WriteFile("./data1",data,0644)
 	if err != nil{
 		panic(err)
	}

	read1,err := ioutil.ReadFile("./data1")
	fmt.Printf("%s",string(read1))

	//相比较于ioutil.WriteFile/ReadFile,os.Create/Open返回的File结构更有灵活性,同时需要注意的是,File结构需要手动关闭,最好是使用defer
	file1,err := os.Create("./data2")
	defer file1.Close()

	bytes,_ := file1.Write(data)
	fmt.Printf("write %d bytes to file\n",bytes)


	file2,err := os.Open("./data2")
	defer file2.Close()


	read2 := make([]byte,len(data))
	bytes,_ = file2.Read(read2)
	fmt.Printf("read %d bytes from file\n",bytes)
	fmt.Println(string(read2))
}
