package chapter5

import (
	"fmt"
	"../golang.org/x/net/html"
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

func ChangeStack(stack []int32,num int32) []int32{
	stack = append(stack,num)
	fmt.Printf("stack:%v\n",stack)
	return stack
}

func Visit(links []string,n *html.Node)[] string{
	if n.Type == html.ElementNode && n.Data == "a"{
		for _,a := range n.Attr{
			if a.Key == "href"{
				links = append(links,a.Val)
			}
		}
	}

	for c :=  n.FirstChild;c != nil;c=c.NextSibling{
		links = Visit(links,c)
	}
	return links
}