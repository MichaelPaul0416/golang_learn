package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"fmt"
)

//介绍golang的二进制读写工具

type postGob struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	//构造一个带缓冲区的编码器
	encoder := gob.NewEncoder(buffer)
	//编码器对数据进行编码
	err := encoder.Encode(data)

	if err != nil {
		panic(err)
	}
	//指定文件名,文件内容以及文件的权限
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	//返回原始数据和一个err
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//将返回的字节slice转换成缓冲区
	buffer := bytes.NewBuffer(raw)
	//构造解码器
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	pg := postGob{Id:1,Content:"jane",Author:"paul"}
	store(pg,"postGob")
	var postRead postGob
	load(&postRead,"postGob")
	fmt.Println(postRead)
}
