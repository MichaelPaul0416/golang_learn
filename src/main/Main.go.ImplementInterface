package main

import (
	"../chapter7"
	"fmt"
	"flag"
	"time"
	"io"
	"os"
	"bytes"
)

func main() {
	//空的接口类型可以赋值给任意的类型
	var any chapter7.Empty
	any = true
	fmt.Printf("empty -> boolean:%t\n", any)
	any = 1
	fmt.Printf("empty -> int:%d\n", any)

	//使用.Value解析
	//flag.Duration：创建一个命令行参数，参数名是period，参数的默认值是1s，也就是第二个参数指定值，参数解释是第三个参数值
	var period = flag.Duration("period", 1*time.Second, "sleep period") //返回的是一个指针类型
	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println("done")

	//返回一个指针类型
	var c chapter7.Celsius
	c.Init(20.0)
	var temp = chapter7.ChangeTemperature("temp", c, "the temperature")
	flag.Parse()
	fmt.Printf("%v\n",*temp)

	//接口值=动态类型[java多态中的实际类型]+动态值[实际类型的实例对象]
	//输出动态类型
	var w io.Writer
	w = os.Stdout
	fmt.Printf("real type:%T\n",w)
	w = new(bytes.Buffer)
	fmt.Printf("real type:%T\n",w)

	//debug := true
	debug := false//使用false话，那么仅仅只是申明了buf的类型，没有关联实例对象[java中的话，一般会报空指针]
	//下面这行语句，对于接口值的赋值，仅仅是给了动态类型是一个byte.Buffer的指针，但是对于动态值，没有进行赋值
	var buf *bytes.Buffer
	fmt.Printf("pointer null:%t\n",buf == nil)//输出true
	if debug{
		buf = new(bytes.Buffer)//赋值了动态类型和动态值，所以接口值是不为空的
	}

	buf.WriteString("hello")
	fmt.Printf("%s\n",buf.String())

}
