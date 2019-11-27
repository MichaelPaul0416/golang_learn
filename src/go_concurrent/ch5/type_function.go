package main

import "fmt"

// http中的HandlerFunc类型，作用是将一个函数转换为实现了目的接口类型的struct
type FunHandler func(n string,age *int32) int32


// Handler 接口(类比http中的)
type SimpleFun interface {
	Display(n string,age *int32) int32;
}

// 有点像代理模式，给你增加一些前置后置的逻辑，用以增强，将一个函数func转换为一个接口类型
func (f FunHandler) Display(n string,age *int32) int32{
	// 先拦截处理一下参数
	fmt.Printf("name:%s\t and age:%d\n",n,*age)
	*age ++
	// 回调传入的函数类型
	return f(n,age)
}

func FuncInstance(n string) (SimpleFun,string){
	// FunHandler实现了SimpleFun接口方法,所以返回一个SimpleFun
	return FunHandler(func(n string,age *int32) int32{
		fmt.Printf("n->%s\n",n)
		return *age
	}),"name:" + n
}
func main(){
	p := func(n string,age *int32) int32{
		fmt.Printf("customer func:%s\t%d\n",n,*age)
		return *age * 2
	}


	num := int32(1)
	p("paul",&num)

	f,n := FuncInstance("George")
	num = int32(1)
	fmt.Printf("final number:%d\n",f.Display(n,&num))

}
