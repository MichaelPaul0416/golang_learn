package chapter5

import (
	"fmt"
	"../golang.org/x/net/html"
	"net/http"
	"io"
	"os"
	"bufio"
)

type NodeType int32

const (
	ErrorNode    int = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

//type Node struct {
//	Type                    NodeType
//	Data                    string
//	Attr                    []Attribute
//	FirstChild, NextSibling *Node
//}

//type Attribute struct {
//	Key, Val string
//}

func ShowConst() {
	fmt.Printf("Document:%d\n", DocumentNode)
}

func ChangeStack(stack []int32, num int32) []int32 {
	stack = append(stack, num)
	fmt.Printf("stack:%v\n", stack)
	return stack
}

func Visit(links []string, n *html.Node) [] string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

func FetchAndParse(url string) {
	read, err := fetch(url)
	if err != nil {
		fmt.Printf("http get error :%s\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(read)
	if err != nil {
		fmt.Printf("html parse error :%s\n", err)
		os.Exit(1)
	}

	a := Visit(nil, doc)
	for _, link := range a {
		fmt.Println(link)
	}
}

// 将http调用迁移到内部,多返回值
func fetch(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %S", url, resp.StatusCode)
	}
	r := resp.Body
	//一般都是要关闭流的，这里没有关闭只是偷懒
	//resp.Body.Close()
	return r, nil
}

func ReadUntilEOF() (string, error) {
	in := bufio.NewReader(os.Stdin)

	var s string
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			return "", fmt.Errorf("read failed:%v", err)
		}

		s += string(r)
	}
	return fmt.Sprintf("line:%v\n", s), nil
}

func Square(n int) int {
	return n * n
}

func Negative(n int) int {
	return -n
}

func Product(m, n int) int {
	return m * n
}

//后面的func 其实就是类似java8里面传了一个lambda表达式进来，一套计算规则作为参数传入
func Composite(m int, c func(n int) int) int {
	return m + c(m)
}

//匿名函数
func WithoutNameFunc(i int) func(o int) int {
	m := i + 1
	return func(n int) int {
		m += 2
		return m * n
	}
}
