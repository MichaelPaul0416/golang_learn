package main

import (
	"../chapter6"
	"fmt"
)
func main(){
	//方法调用
	p := chapter6.Point{1,3}
	q := chapter6.Point{2,4}

	//绑定到类型上，而不是绑定到指针上
	p.MovePointer(1)
	fmt.Printf("move point:%v\n",p)//绑定的对象不是指针类型，所以即便MovePointer内部对对象做了修改，但是还是不会改变原先的

	//将q作为一个参数，传入给p的方法Distance，进行计算调用
	fmt.Printf("distance:%.2f\n",p.Distance(q))

	cells := chapter6.Path{
		{1,1},
		{1,2},
		{2,1},
		{1,1},
	}

	fmt.Printf("三角形周长:%.2f\n",cells.Distance())

	//指针接收者
	point := &chapter6.Point{1,3}
	point.MovePoint(1)//{2,4}
	//这样可以的，因为实参接收者point本身就是一个指针类型的变量，可以获取地址
	(*point).MovePoint(2)//{4,6}
	fmt.Printf("move:%v\n",*point)

	//下面这种声明也是可以的，会对变量进行隐式转换（&point）
	//point := chapter6.Point{1,3}


	//不能对一个不能取地址的接收者参数调用指针接收者方法，因为这无法取得临时变量的地址
	//chapter6.Point{1,3}.MovePoint(1)
	//但是这一行代码可以，因为通过&显示获取了变量的地址
	(&chapter6.Point{1,3}).MovePoint(2)

	head := chapter6.IntList{Value:1,Next:nil}
	var node_1 chapter6.IntList
	node_1.Value = 2
	head.Next = &node_1
	node_1.Next = &chapter6.IntList{Value:3,Next:nil}
	fmt.Printf("linked list sum:%d\n",head.SumList())

	var ml chapter6.MapList
	ml = chapter6.MapList{"two":{11,12}}//初始化对象
	ml.Put("one",1)
	ml.Put("one",2)

	fmt.Printf("sum list of key:%d\n",ml.Get("two"))
	fmt.Printf("mapList:%v\n",ml)

}
