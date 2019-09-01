package main

import (
	"os"
	"encoding/csv"
	"strconv"
	"fmt"
)

//使用csv进行文件的写入和读取
type post struct {
	Id      int
	Content string
	Author  string
}

func main(){
	csvFile,err := os.Create("./post.csv")
	if err != nil{
		panic(err)
	}

	defer csvFile.Close()
	allPosts := []post{
		post{Id:1,Content:"Hello world",Author:"Michael"},
		post{Id:2,Content:"golang web programming",Author:"George"},
		post{Id:3,Content:"Java concurrency coding",Author:"Paul"},
		post{Id:4,Content:"java distribute system",Author:"Jane"},
	}

	//将一个文件包装为一个csv的流
	writer := csv.NewWriter(csvFile)
	for _,p := range allPosts{
		line := []string{strconv.Itoa(p.Id),p.Content,p.Author}
		err := writer.Write(line)
		if err != nil{
			panic(err)
		}
	}

	writer.Flush()
	fmt.Printf("csv write done...\n")

	file,err := os.Open("./post.csv")
	if err != nil{
		panic(err)
	}

	defer file.Close()
	//包装为一个csv的输入流
	reader := csv.NewReader(file)
	//读取器在读取时发现记录里面缺少了某些字段,读取的进程也不会中断
	//如果这个值设置为正数,就是用户要求从每条记录里面读取出的字段数量,并且当读取器从csv文件里里面读取出的字段数量少于这个值的时候,就会抛出一个错误
	//如果这个字段设置为0,那么读取器就会将读取渠道的第一条记录的字段数量作为它的值
	reader.FieldsPerRecord = -1
	record,err := reader.ReadAll()
	if err != nil{
		panic(err)
	}

	var posts []post
	for _,item := range record{
		id ,_ := strconv.ParseInt(item[0],0,0)
		p := post{Id:int(id),Content:item[1],Author:item[2]}
		posts = append(posts,p)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}
