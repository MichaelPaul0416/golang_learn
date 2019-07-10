package chapter5

import (
	"fmt"
	"runtime/debug"
)

//处理宕机(异常)+延迟函数defer

func Div(a, b float64) float64 {

	defer debug.PrintStack()
	if b == 0 {
		panic("被除数不能为0")
	}
	//下面语句的注释打开的话，不会被执行，因为panic已经抛出异常
	//defer debug.PrintStack()
	fmt.Printf("start to div[%f/%f\n]",a,b)
	return float64(a / b)
}


func FuncWithDeferAndPanic(i int){
	//如果一个函数内部有多个defer，并且多个defer之间没有panic抛出宕机错误，那么defer的执行顺序是按照defer由下到上的逆序输出

	defer func() {
		fmt.Printf("first defer output:%d\n",i)
	}()

	fmt.Printf("input number:%d\n",i)

	defer func() {
		fmt.Printf("second defer output:%d\n",i)
	}()

	//如果打开下面的注释，那么最终会输出input number，然后再是second defer,first defer;third defer是不会被输出
	//panic("inner panic")

	defer func() {
		fmt.Printf("third defer output:%d\n",i)
	}()
}

type Protocol int

const (
	HTTP Protocol = iota
	RPC
	JAR
	FTP
	SFTP
	DNS
)

//recover恢复宕机panic
func DealPartPanic(proto int)(pro Protocol,err error){
	//自定义的错误类型
	type bailout struct {}

	defer func() {
		//recover()函数捕获到抛出的panic，根据panic的具体类型决定下一步的操作
		switch p := recover();p {
		case nil:
			fmt.Printf("none panic occured\n")
		case bailout{}:
			err = fmt.Errorf("current protocol invalid")
		default:
			//不再当前函数处理范围之内，继续panic
			panic(p)
		}
	}()

	if proto < 0 || proto > 6{
		panic("wrong type")
	}

	if proto == 3{
		panic(bailout{})
	}

	return Protocol(proto),nil
}