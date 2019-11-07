package main

import (
	"fmt"
	"sync/atomic"
)
type sim int

func main(){
	var i int32 = 10
	fmt.Printf("number(+3):%d\n",atomic.AddInt32(&i,3))

	var m uint32 = 10
	// 对于uint类型的数据，如果要实现减法的话，先将减数N（N>0）-1 暨N-1
	// 然后对其取反^N，其实就是获取减数N的相反数(负数)表示形式
	// 3 -> 0011
	// -3 -> 1101
	// 最后利用补码的形式进行加减运算
	fmt.Printf("number(-3):%d\n",atomic.AddUint32(&m,^uint32(3-1)))

	ok := atomic.CompareAndSwapInt32(&i,13,20)
	if ok{
		fmt.Printf("cas success:%d\n",i)
	}else{
		fmt.Printf("cas failed,and old value:%d\n",i)
	}

	for{
		// 原子性的读取并且赋值给v
		v := atomic.LoadInt32(&i)
		if atomic.CompareAndSwapInt32(&i,20,16);ok{
			fmt.Printf("cas success:old(%d) -> new(%d)\n",v,i)
			break
		}
	}

	// swap
	// 与CAS不同的是，swap不会关心旧值，直接设置新值
	old := atomic.SwapInt32(&i,9)
	fmt.Printf("swap:old(%d)/new(%d)\n",old,i)

	var atomicVal atomic.Value
	g := sim(1)
	atomicVal.Store(g)
	fmt.Printf("atomic value:%d\n",atomicVal.Load())


	// 将atomic.Value作为参数传递给方法修改，不会影响原来的atomic.Value以及原先Store的值，因为操作的都是副本
	var countVal atomic.Value
	countVal.Store([]int{1,3,5,7})
	anotherStore(countVal)
	fmt.Printf("count atomic.Value:%v\n",countVal.Load())

	// 下面这种直接将Store的入参作为参数传递给别的方法进行修改，会影响Load的结果
	var cv atomic.Value
	pm := []int{1,2,3}
	cv.Store(pm)
	anotherStoreValue(pm)
	fmt.Printf("count value：%v\n",cv.Load())

}

func anotherStore(countVal atomic.Value){
	countVal.Store([]int{2,4,6,8})
}


func anotherStoreValue(m []int){
	m[0] = 9
}