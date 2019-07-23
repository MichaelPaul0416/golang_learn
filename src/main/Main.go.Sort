package main

import (
	"../chapter7"
	"sort"
	"fmt"
)

func main(){
	s1 := make([]string,10)
	s1 = append(s1,"b")
	s1 = append(s1,"c")
	chapter7.SortSlice(s1)

	s1 = append(s1,"a")
	sort.Strings(s1)
	fmt.Printf("go api sort:%v\n",s1)

	chapter7.SortTracksByArtist()

	chapter7.SortWithMulti()

	var ary = []int{2,3,1,5,4,1}
	fmt.Printf("is sorted:%t\n",sort.IntsAreSorted(ary))//sort.IntsAreSorted不会去调用sort#Swap方法，只调用其他两个
	sort.Ints(ary)
	fmt.Printf("is sorted:%t\n",sort.IntsAreSorted(ary))

	is := sort.IntSlice(ary)//封装，实现sort接口的方法
	fmt.Printf("ary:%v\n",is)
	//sort.Reverse只接受sort接口的实现，所以slice不能直接传入，需要使用sort提供的IntSlice封装一下
	sort.Sort(sort.Reverse(is))
	fmt.Printf("reverse ary:%v\n",is)
	//因为这个方法最终是去调用对应类的Less方法，而reverse只是将原先传入的i,j顺序掉换成j,i
	//原先的Less返回的是排好序之后s[i] < s[j] = true
	//现在i->j j->i, 所以s[j] < s[i] = false
	//所以此时IsSorted返回false
	fmt.Printf("is sorted:%t\n",sort.IsSorted(is))
}
