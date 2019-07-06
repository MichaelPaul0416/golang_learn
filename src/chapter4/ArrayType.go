package chapter4

import (
	"fmt"
	"crypto/sha256"
)

func DefArray() {
	var a1 = [3]int{1, 2, 3}

	fmt.Printf("定长数组定义：%v\n", a1)

	var a2 = [...]string{"hello", "world", "go"}
	fmt.Printf("变长数组定义：%v\n", a2)

	var a3 = [...]int{10: 2, -1}
	fmt.Printf("指定长度但是没有指定下标值：%v\n", a3)

	var a4 = [...]int{1, 2, 3}
	fmt.Printf("比较a1、a3：%t\n", a1 == a4)
	//数组的比较只能比较长度相同的两个数组
	//fmt.Printf("a1和a3的比较会编译错误：%v\n",a1 == a3)

	b1 := sha256.Sum256([]byte("x"))
	b2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", b1, b2, b1 == b2, b1)

}

func ChangeArray(a [3]int,index int){

	if index >= len(a){
		fmt.Printf("修改下标不合法\n")
		return
	}

	a[index] = -1
	fmt.Printf("方法内数组：%v\n",a)
}

func ChangeArrayByPoint(a *[3]int,index int){
	if len(*a) <= index{
		fmt.Printf("修改下标不合法\n")
		return
	}

	a[index] = -1
	fmt.Printf("使用指针，修改数组，方法内数组：%v\n",*a)
}
