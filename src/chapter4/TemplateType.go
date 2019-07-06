package chapter4

import (
	"html/template"
	"os"
	"log"
)

type Person struct {
	Name string
	Books []Book
}

type Book struct {
	BookName string
	Price    float64
	Author   string
	BookId   int64
}

var books = [3]Book{
	{"Java", 12.4, "Paul", 1},
	{"Linux", 34.2, "Linus", 2},
	{"Golang", 23.5, "George", 3},
}

//{{rang .Books}}中的.Books表示Person对象中的成员变量Books
const temp = `{{.Name}}'s reading books:
{{range .Books}}------------------------
BookName: {{.BookName}}
Price: {{.Price | printf "%.2f $"}}
Author: {{.Author | authorInfo}}
{{end}}`

func authorInfo(s string) string {
	return s + "--China"
}

var report = template.Must(template.New("bookList").Funcs(template.FuncMap{"authorInfo": authorInfo}).Parse(temp))

func GenerateFromTemplate(){
	var person = new(Person)
	person.Name = "Michael"
	person.Books = books[:]

	if err := report.Execute(os.Stdout,person); err != nil{
		log.Fatal(err)
	}
}
