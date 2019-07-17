package chapter6

import (
	"image/color"
	"fmt"
	"math"
	"sync"
)

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type SharedMessage struct {
	*Point
	Desc string
}

type PointPath []Point

func (p PointPath) TranslateBy(offset Point, add bool) {
	//对p，q两个Point操作，返回一个新的Point
	var op func(p, q Point) Point

	if add {
		fmt.Println("call add")
		op = Point.Add
	} else {
		fmt.Println("call sub")
		op = Point.Sub
	}

	//记录
	for i:= range p{
		p[i] = op(p[i],offset)
	}

	fmt.Printf("len:%d\tcap:%d\n",len(p),cap(p))

}

//定义一个匿名的结构体，并且初始化
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	//初始化
	mapping: make(map[string]string),
}

func Put(key, val string) {
	cache.Lock()
	cache.mapping[key] = val
	cache.Unlock()
}

func Get(key string) string {
	//直接调用匿名成员变量sync.Mutex的Lock/Unlock方法
	cache.Lock()
	val := cache.mapping[key]
	cache.Unlock()
	return val
}

func (cp *SharedMessage) SetPoint(p *Point) {
	cp.Point = p

}

func (cp SharedMessage) ToString() {
	fmt.Printf("ShareMessage[Point:%v\tDesc:%s]\n", *cp.Point, cp.Desc)
}

//当结构体中绑定方法A，结构体的匿名成员变量P也有同名方法A[方法入参和出参都不影响]
//那么会先寻找结构体中的方法，如果找到，直接调用，如果没找到，再继续调用匿名成员变量P的方法
func (cp ColoredPoint) Distance(q Point) float64 {
	fmt.Println("choose ColoredPoint method")
	return math.Hypot(cp.X-q.X, cp.Y-q.Y)
}

func Init(x float64) ColoredPoint {
	var cp ColoredPoint
	cp.X = x
	cp.Point.Y = x + 1
	return cp
}

func (cp *ColoredPoint) ChangePoint(x, y float64) {
	cp.X = x
	cp.Y = y
}
