package chapter6

import (
	"math"
)

type Point struct {
	X, Y float64
}

//方法，通过(p Point)将Distance函数绑定到Point类型上
//p Point:称为方法的接收者
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//包级别对外公开的函数
func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

//可以为任何类型添加方法，除非它的类型既不是指针类型也不是接口类型
//slice类型可以被添加方法
//同时方法可以同名，因为他们有不同的命名空间，分属于不同给的类型
//类型所拥有的方法名必须是唯一的，但是不同的类型可以使用相同的方法名
func (pa Path) Distance() float64 {
	sum := 0.0

	//方法内部可以直接调用方法所绑定的对象(pa)
	for index := range pa {
		if index > 0 {
			sum += pa[index].Distance(pa[index-1])
		}
	}
	return sum
}

//指针接受者的方法
func (p *Point) MovePoint(offset float64) {
	p.X += offset
	p.Y += offset
}

//nil作为方法接收者
type IntList struct {
	Value int
	Next  *IntList
}

func (list *IntList) SumList() int{
	if list == nil{
		return 0
	}
	//递归调用
	return list.Value + list.Next.SumList()
}

type MapList map[string][]int

//返回key=k对应的list的sum
func (m MapList) Get(k string) int{
	if l := m[k]; l != nil{
		sum := 0
		for _,v := range l{
			sum += v
		}
		return sum
	}

	return 0
}

func (m MapList) Put(k string,v int) {
	if m[k] != nil{
		l := m[k]
		l = append(l,v)
		m[k] = l
	}else {
		t := make([]int,1)
		t = append(t,v)
		m[k] = t
	}
}
