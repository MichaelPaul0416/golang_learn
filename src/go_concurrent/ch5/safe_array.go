package main

import (
	"sync/atomic"
	"errors"
)

type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)

	Get(index uint32) (elem int, err error)

	Len() uint32
}

type concurrentArray struct {
	length uint32
	val    atomic.Value
}

func NewConcurrentArray(length uint32) ConcurrentArray {
	array := concurrentArray{}
	array.length = length
	// 设置val为长度不变的slice
	array.val.Store(make([]int, array.length))
	return &array
}

func (array *concurrentArray) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index);err != nil{
		return err
	}

	newArray := make([]int,array.length)
	// 类型断定，是一个[]int 不然就会引发一个panic
	copy(newArray,array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return nil
}

func (array *concurrentArray) Get(index uint32)(elem int,err error){
	if err = array.checkIndex(index);err != nil{
		return
	}

	// 类型断言
	elem = array.val.Load().([]int)[index]
	return
}

func (array *concurrentArray) Len() uint32{
	return array.length
}

func (array *concurrentArray) checkIndex(index uint32) (error) {
	if index > array.length-1{
		return errors.New("index illegal")
	}
	return nil
}
