package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//处理一对多-多对一的关系,基础api

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}

type Article struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Article *Article
}

func (comment *Comment) Create() (err error) {
	if comment.Article == nil {
		err = errors.New("article not found")
		return
	}

	stmt,err := Db.Prepare("insert into Comments (content,author,post_id) values(?,?,?)")
	if err != nil{
		fmt.Printf("err:%v\n",err)
		return
	}
	res,err := stmt.Exec(comment.Content, comment.Author, comment.Article.Id)
	if err != nil{
		fmt.Printf("err:%v\n",err)
		return
	}

	id,err := res.LastInsertId()
	comment.Id = int(id)
	return
}

func GetArticle(id int) (article Article, err error) {
	article = Article{}
	article.Comments = []Comment{}

	err = Db.QueryRow("select id,content,author from Article where id ?", id).Scan(
		&article.Id, &article.Content, &article.Author)

	rows, err := Db.Query("select id,content,author from Comments")
	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{Article: &article}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}

		article.Comments = append(article.Comments, comment)
	}

	rows.Close()
	return
}

func (article *Article) Create() (err error) {
	stmt,err := Db.Prepare("insert into Article (content,author) values(?,?)")
	if err != nil{
		fmt.Printf("err:%v\n",err)
		return
	}
	res,err := stmt.Exec(article.Content,article.Author)
	if err != nil{
		fmt.Printf("err:%s\n",err)
		return
	}
	id,_ := res.LastInsertId()
	defer stmt.Close()
	article.Id = int(id)
	return
}
func main() {
	a := Article{Content: "Hello world!", Author: "Jane"}
	a.Create()

	comment := Comment{Content:"Goo post!",Author:"Jane",Article:&a}
	comment.Create()

	readPost,_ := GetArticle(a.Id)

	fmt.Println(readPost)
}
