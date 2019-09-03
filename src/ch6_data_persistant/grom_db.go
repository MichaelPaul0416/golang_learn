package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
	"math/rand"
)


type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
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
	//不设置这个参数，会把表名转义后加s
	dbGrom.SingularTable(true)

	//如果数据库里面没有对应的表，则根据结构体生成对应的表

	dbGrom.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&Like{})
}

func main() {

	like := &Like{
		Ip:        "127.0.0.1",
		Ua:        "aaa",
		Title:     "bbb",
		Hash:      uint64(rand.Intn(1000)),
		CreatedAt: time.Now(),
	}
	if err := dbGrom.Create(like).Error; err != nil {
		fmt.Printf("err:%v\n",err)
	}

	fmt.Printf("after insert:%v\n",like)

}
