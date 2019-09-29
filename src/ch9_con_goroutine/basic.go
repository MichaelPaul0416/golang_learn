package main

import "fmt"

func main(){
	fmt.Printf("%s\n",`1+2`)

	fmt.Printf("%s\n","\a")

	s,r := setValueWhileReturn(3)
	fmt.Printf("%d\t%s\n",s,r)
}


/**
当函数有返回值的时候，可以在函数体内部直接使用返回的函数值
给函数值赋值就相当于给返回值赋值
 */
func setValueWhileReturn(m int)(s int,r string){
	fmt.Printf("input arg:%d\n",m)
	r = "hello world"
	s = 3

	// 如果函数声明的结果是有名称的，那么return关键字后面就不需要在追加任何东西
	return
}