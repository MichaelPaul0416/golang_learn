package chapter7

import (
	"io"
	"os"
	"fmt"
)

func AssertTip(){
	var w io.Writer
	w = os.Stdout
	//w.Close()//此时w声明的类型为io.Writer,不具备Close方法

	f,ok := w.(*os.File)//断言w为*os.File类型-->true
	fmt.Printf("assert io.Writer --> *os.File:%t\n",ok)
	f.Close()//返回的f是断言的类型，也就是*os.File,具有Close方法
}
