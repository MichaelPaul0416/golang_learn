package main

import (
	"runtime/debug"
	"sync/atomic"
	"sync"
	"fmt"
	"runtime"
)

func main(){
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count,1)
	}
	pool := sync.Pool{New:newFunc}

	v1 := pool.Get()
	fmt.Printf("get obj from pool:%v\n",v1)

	// 将对象放入到池中，以供后序获取使用
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	// 从池中随机获取一个，优先从通过Put方法放入的对象中获取
	v2 := pool.Get()
	fmt.Printf("get v2 from pool:%v\n",v2)
	v2 = pool.Get()
	fmt.Printf("get v2 from pool again:%v\n",v2)
	v2 = pool.Get()
	fmt.Printf("get v2 from pool again:%v\n",v2)
	// 通过Put放入的获取完了，继续获取的话，就通过New字段指定的函数来生成对象，但是这个对象不会缓存在pool中，生成之后会立即返回给调用方
	v2 = pool.Get()
	fmt.Printf("get v2 from pool again:%v\n",v2)

	// 开始gc
	// gc后对象池中的所有缓存对象（通过Put方法放入的，都会被回收，所以下次开始获取的话，就是直接从New字段指定的func中获取对象）
	debug.SetGCPercent(100)
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("get v3 from pool after gc:%v\n",v3)
	// 将New对应的func也置为空了之后，此时再调用Get方法，返回的就是nil
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("get v4 from pool after set nil to New field:%v\n",v4)
}