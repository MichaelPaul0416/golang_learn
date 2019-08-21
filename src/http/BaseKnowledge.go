package http

import (
	"math"
	"fmt"
	"reflect"
	"runtime"
)

func SumNumbers(a, b int) int {
	return a + b
}

func QueryThirdSide(a, b float64) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func ShowSlice(s []int) {
	if s == nil {
		fmt.Errorf("empty slice")
		return
	}

	for i := range s {
		fmt.Printf("%d\t", i)
	}
	fmt.Printf("\n")
}

func CleanZeroByte() {
	const (
		read    = 1 << iota
		write
		execute
	)

	access := 7
	fmt.Printf("read:%t\twrite:%t\texecute:%t\n", access&read == read, access&write == write, access&execute == execute)

	fmt.Printf("clean execute\n")
	/**
	&^：清除某个位上的数字，当这个操作符右边是1的时候，左边无论是什么，结果都是0，当操作符右边是0的时候，结果就是左边的值
	 */
	access = access &^ execute
	fmt.Printf("read:%t\twrite:%t\texecute:%t\n", access&read == read, access&write == write, access&execute == execute)
}

func Apply(op func(int,int) int,a,b int) int{
	p := reflect.ValueOf(op).Pointer()

	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("calling function %s with args:%d,%d\n",opName,a,b)
	return op(a,b)
}