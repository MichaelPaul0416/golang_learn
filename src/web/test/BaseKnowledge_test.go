package test

import (
	"testing"
	//绝对路径导包，相当于$GOPATH/http，同时$GOPATH一般就是src目录的上一层
	web "web/example"
	"fmt"
	"strings"
)

func TestSumNumbers(t *testing.T) {
	if i := web.SumNumbers(1, 2); i != 3 {
		t.Error("error while call SumNumber")
	} else {
		t.Log("ok")
	}
}

func TestQueryThirdSide(t *testing.T) {
	if i := web.QueryThirdSide(3, 4); i < 0 {
		t.Error("third side error")
	} else {
		t.Logf("third side is:%.2f\n", i)
	}
}

func TestShowSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	web.ShowSlice(s)
}

func TestCleanZeroByte(t *testing.T) {
	web.CleanZeroByte()
}

func TestApply(t *testing.T) {
	f := func(a, b int) int {
		return a + b*a
	}
	r := web.Apply(f,1,2)
	fmt.Printf("return :%d\n",r)
}

func TestSlice(t *testing.T){
	p := []int{1,2,3,4,5}
	p = append(p,3)
	fmt.Printf("show slice:%v\n",p)
}

func TestStringReplace(t *testing.T){
	times := "2019-08-26T16:38:19Z"
	fmt.Println(strings.Replace(times,"T"," ",1))

}
