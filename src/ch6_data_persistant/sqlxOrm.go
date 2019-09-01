package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

//使用orm框架-sqlx

type ArticleSqlx struct {
	Id int
	Content string
	//对应数据库中的表的字段名
	AuthorName string `db:author`
}

var SqlXDb *sqlx.DB
func init(){
	var err error
	//SqlXDb ,err = sqlx.Open("mysql","user=root dbname=golang password=root charset=utf8 parseTime=true")
	SqlXDb ,err = sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=true")
	if err != nil{
		panic(err)
	}
	fmt.Println("sqlx connect successfully")
}

func GetArticleSqlX(id int)(a ArticleSqlx,err error){
	a = ArticleSqlx{}
	err = SqlXDb.QueryRowx("select id,content,author from Article where id=?",id).StructScan(&a)
	if err != nil{
		fmt.Printf("err:%v\n",err)
		return
	}
	return
}

func main(){
	a,_ := GetArticleSqlX(1)
	fmt.Println(a)
}