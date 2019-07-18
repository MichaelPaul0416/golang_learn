package main

import (
	"../chapter6"
	"fmt"
)

func main(){

	is := chapter6.IntSet{make([]uint64,0)}
	is.Add(69)
	is.Add(128)
	is.Add(1)
	fmt.Printf("is:%v\n",is)
	fmt.Printf("contain 100 : %t\n",is.Has(100))

	temp := chapter6.IntSet{make([]uint64,0)}
	temp.Add(67)
	temp.Add(24)
	is.UnionWith(&temp)
	fmt.Printf("new Inset:%v\n",is.Words)
	fmt.Printf("contains %d:%t\n",24,is.Has(24))

	fmt.Printf("slice:%v\n",is.String())

	cp := chapter6.IntSet{make([]uint64,0)}
	cp.Add(67)
	fmt.Printf("big:%v\n",cp.Words)

	co := chapter6.Count{}
	co.Init(65,"one")
	fmt.Printf("Count:%v\n",co)
	co.Add()
	fmt.Printf("Add Count:%v\n",co)
}