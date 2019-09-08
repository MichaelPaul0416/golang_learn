package main

import (
	"testing"
	"time"
	"fmt"
)

/**
对go的功能性测试,所有的测试函数要以Test开头,并且入参是*testing.T
 */

/**
使用go test -v -cover -short -parallel 3 进行单元测试
-parallel 3 指定了并行测试的数目,也就是说下面三个TestParallel_x的单元测试可以并行测试
-cover 覆盖率标志,获取测试用例对代码的覆盖标志 当测试用例中有用到项目的方法或者函数的时候,这个参数才可以使用
-short 当某个单元测试耗费的时间比较长的时候,可以在代码内通过t.Skipped()判断是否需要跳过这次的单元测试
 */
func TestParallel_1(t *testing.T){
	t.Parallel()
	time.Sleep(time.Second * 1)
	fmt.Printf("worker-1 done\n")
}

func TestParallel_2(t *testing.T){
	t.Parallel()
	time.Sleep(time .Second * 2)
	fmt.Printf("worker-2 done\n")
}

func TestParallel_3(t *testing.T){
	t.Parallel()
	time.Sleep(time.Second * 3)
	fmt.Printf("worker-3 done\n")
}
