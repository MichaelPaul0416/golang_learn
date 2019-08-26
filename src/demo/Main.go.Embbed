package main

import (
	"../chapter6"
	"fmt"
	"image/color"
)

func main(){

	cp := chapter6.Init(0.3)
	fmt.Printf("ColoredPoint:%v\n",cp)

	//调用成员变量Point的方法
	red := color.RGBA{255,0,0,255}
	blue := color.RGBA{0,0,255,255}
	var p = chapter6.ColoredPoint{chapter6.Point{1.1,2.2},red}
	var q = chapter6.ColoredPoint{chapter6.Point{1.1,2.3},blue}
	//Point的方法都被归入到ColoredPoint中，可以直接调用
	//而且Point作为成员变量，在ColoredPoint中必须是直接被声明，不能有别名
	fmt.Printf("distance[p->q]:%.2f\n",p.Distance(q.Point))

	//传入指针，修改匿名成员的属性的方法
	(&cp).ChangePoint(1.2,3.4)
	//下面这样也可以
	//cp.ChangePoint(1.2,3.4)
	fmt.Printf("changed ColoredPoint:%v\n",cp)

	//匿名成员变量指针
	sharePoint := chapter6.Point{1,2}
	sm1 := chapter6.SharedMessage{&sharePoint,"hello"}
	fmt.Printf("sm1:%v\tdesc:%s\n",*sm1.Point,sm1.Desc)
	sm2 := chapter6.SharedMessage{&chapter6.Point{2,3},"world"}
	fmt.Printf("sm2:%v\t%s\n",*sm2.Point,sm2.Desc)
	//sm2与sm1共享
	sm2.SetPoint(sm1.Point)
	fmt.Printf("after set:%v\t%s\n",sm2.Point,sm2.Desc)

	//封装独占锁，使用缓存
	chapter6.Put("j","java")
	fmt.Printf("cache value:%s\n",chapter6.Get("j"))

	//方法作为一个表达式
	p1 := chapter6.Point{1,2}
	p2 := chapter6.Point{4,6}

	//将方法赋值给变量
	distanceFromP1 := p1.Distance
	fmt.Printf("distance from p1 to p2:%.2f\n",distanceFromP1(p2))
	fmt.Printf("distance from p1 to (0,0):%.2f\n",distanceFromP1(chapter6.Point{}))

	//2s之后计算
	//由另外一个线程执行，不是main线程，需要打开下面的for死循环才可以打印
	//time.AfterFunc(2 * time.Second,sm1.ToString)
	//for ; ;  {
	//	fmt.Println("00")
	//	time.Sleep(1000 * time.Second)
	//}

	pa := chapter6.Point{2,3}
	pp := chapter6.PointPath{pa}
	pp.TranslateBy(chapter6.Point{1,3},true)
	pp.TranslateBy(chapter6.Point{1,1},true)
	fmt.Printf("point path:%v\n",pp)
}
