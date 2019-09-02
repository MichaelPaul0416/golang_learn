package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
)

type art struct {
	id       int
	content  string
	author   string `sql:"not null"`
	comments []comm
	created  time.Time
}

type comm struct {
	id      int
	content string
	author  string `sql:"not null"`
	artId   int    `sql:"index"`
	created time.Time
}

var dbGrom *gorm.DB

func init() {
	var err error
	dbGrom, err = gorm.Open("mysql", "root:root@/golang?charset=utf8&parseTime=true")
	if err != nil {
		panic("connect to db failed")
	}else{
		fmt.Printf("connect to db successfully\n")
	}

	//如果数据库里面没有对应的表，则根据结构体生成对应的表
	dbGrom.AutoMigrate(&art{}, &comm{})
}

func main() {
	a := art{content: "hello", author: "jane",comments:[]comm{}}
	fmt.Println(a)

	dbGrom.Create(&a)
	fmt.Printf("after insert:%v\n", a)

	c := comm{content: "world", author: "jane"}
	//目前下面执行到Association的时候报了空指针，可能原因是Model方法返回的记录为空[因为目前值执行dbGrom.Create的时候，记录没有插到数据库里面]
	dbGrom.Model(&a).Association("comments").Append(c)

	var readArt art
	dbGrom.Where("author = ?", "jane").First(&readArt)

	var cs []comm
	dbGrom.Model(&readArt).Related(&cs)
	fmt.Println(cs)
}
