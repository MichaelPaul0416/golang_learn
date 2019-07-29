package main

import (
	"../chapter7"
	"fmt"
)

func main() {
	chapter7.AssertTip(false)

	chapter7.ShowErrorType()

	chapter7.AssertWithError()

	//定义一个具体的子类Dubbo
	dubbo := chapter7.Dubbo{Proto: "dubbo", Code: 1 << 2}
	chapter7.DynamicTypeByAssert(dubbo)

	http := chapter7.Restful{Proto:"restful",Code: 1 << 1}
	chapter7.DynamicTypeByAssert(http)

	//输出具体的类型
	fmt.Printf("real type:%s\n",chapter7.RealType(true))
}
