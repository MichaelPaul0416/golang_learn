package chapter4

import (
	"time"
	"fmt"
)

type Employee struct{
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}
var coder Employee

type Tree struct{
	value int
	left,right *Tree
}

type Scale struct {
	X int
	Y int
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	//P Point和下面直接使用Point声明等价
	Point//匿名成员
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
	//自定义成员中，如果和匿名成员中有名字一样的变量，那么赋值时会复制给自定义的，而不是匿名成员变量中的同名变量
	Radius int
}

func ShowWheel(){
	var w Wheel
	w.X = 1
	w.Y = 1
	//优先复制给Wheel的变量
	w.Radius = 3
	//需要手动制定匿名变量，然后再复制给匿名变量的Radius
	w.Circle.Radius = 2
	w.Spokes = 5
	//#%v-->格式化输出，连带fieldName一起输出
	fmt.Printf("%#v\n",w)
}

func LargerScale(s Scale,l int) Scale{
	//相当与新建了一个struct并返回,并不是基于原先的修改
	return Scale{s.X * l,s.Y * l}
}

func LargerScaleByPoint(s *Scale,l int){
	s.X *= l
	s.Y *= l
}
func PrintTreeBefore(tree *Tree){
	if tree.left != nil{
		PrintTreeBefore(tree.left)
	}
	fmt.Printf("%d\t",tree.value)

	if tree.right != nil {
		PrintTreeBefore(tree.right)
	}
}
func Sort(a []int) *Tree{
	var root *Tree
	for _,value := range a{
		root = add(root,value)
	}
	return root
}

func add(tree *Tree,value int) *Tree{
	if tree == nil{
		t := new(Tree)
		t.value = value
		return t
	}

	if value < tree.value{
		tree.left = add(tree.left,value)
	}else {
		tree.right = add(tree.right,value)
	}
	return tree
}
func Describe(employee *Employee){
	employee.Salary = 30000
	//作用和上面那行代码一样
	//(*employee).Salary = 30000
	fmt.Printf("employee:%v\n",employee)
}
