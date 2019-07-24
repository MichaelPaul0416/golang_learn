package chapter7

import (
	"io"
	"os"
	"fmt"
)

func AssertTip(c bool) {
	var w io.Writer
	w = os.Stdout
	//w.Close()//此时w声明的类型为io.Writer,不具备Close方法

	f, ok := w.(*os.File) //断言w为*os.File类型-->true
	fmt.Printf("assert io.Writer --> *os.File:%t\n", ok)
	if c {
		f.Close() //返回的f是断言的类型，也就是*os.File,具有Close方法
	}
}

func ShowErrorType() {
	_, err := os.Open("/hello/world")
	fmt.Printf("%v\n", err)  //输出错误信息
	fmt.Printf("%#v\n", err) //输出原始信息
}

//使用断言来返回错误
func AssertWithError() {
	_,err := os.Open("/no/such/file")

	if pe,ok := err.(*os.PathError);ok{
		//如果类型断言成功，那么err被转换为pe的实际类型，也就是断言类型PathError
		fmt.Printf("assert err is *os.PathError:%t\n",ok)
		err = pe.Err
	}
	fmt.Printf("%s\n",err)
}
