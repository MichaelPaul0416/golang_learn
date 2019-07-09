package chapter5

import (
	"fmt"
	"../golang.org/x/net/html"
	"net/http"
	"io"
	"os"
	"bufio"
	"time"
	"log"
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
//函数变量，函数作为一个参数传入
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

//变长函数
func MultiParamFunc(num ...int) {
	var sum = 0
	for _, val := range num {
		sum += val
	}
	fmt.Printf("sum:%d\n", sum)
}

//使用变长函数格式化日志输出
func LogFormatOfError(line int, prefix string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "line:%d\t", line)
	fmt.Fprintf(os.Stderr, prefix, args...)
	fmt.Fprintln(os.Stderr)
}

//defer机制
func FuncWithDefer(num int, tip string) {
	fmt.Printf("hello:%s;this is the %d times;\n", tip, num)
	defer func() {
		fmt.Printf("do finally,closing resource; eg...\n")
	}()
	fmt.Printf("show my welcome\n")
}

func FuncCostTime() {
	//最后加上()代表trace的返回函数会在FuncCostTime方法执行完毕之后被调用
	defer trace("big slow operation")()
	time.Sleep(1 * time.Second)
	fmt.Printf("hello world\n")

}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter func and param is :%s\n", msg)
	return func() {
		log.Printf("exit func and cost:%s\n", time.Since(start))
	}
}

func AbsOfNumber(num, del int) (res int) {
	var result = num - del
	defer func() {
		//里面的内容就是被延迟执行的内容
		if res < 0 {
			res = -res
		}
	}() //最后的这对括号代表需要被延迟至方法结束再去执行定义的匿名函数
	return result
}

func DeferWhileCycle(ns []int) {
	for _, v := range ns {
		fmt.Printf("%d\t", v)
		if v < 0 {
			return
		}
		defer func() {
			//这里输出的v永远都是ns的最后一个元素，或者return时的那个v
			//输出的次数是到目前为止迭代的次数
			fmt.Printf("done:%d\n", v)
		}()
	}
}
