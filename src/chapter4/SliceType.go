package chapter4

import "fmt"

//下面的数组中的数字代表的含义是索引
var numbers = [...]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "night", 10: "ten"}

func ShowNumberMap() {
	fmt.Printf("number-->:%v\n", numbers)
}

func ShowSliceFromArray(start int, end int) {
	if end <= start {
		fmt.Printf("end[%d must >= start[%d]", end, start)
		return
	}

	ary := numbers[start:end]
	fmt.Printf("slice[ary]:[%d:%d] --> %v\n", start, end, ary)
	fmt.Printf("length:%v --> %d\n", ary, len(ary))
	fmt.Printf("capcity:%v --> %d\n", ary, cap(ary))

	var last int = -1
	if end+3 > len(numbers) {
		last = 10
	} else {
		last = end + 3
	}
	arycpy := ary[:last]
	fmt.Printf("arycpy(slice of ary):[%d->%d]\t%v\n", 0, last, arycpy)
	//fmt.Printf("")
}

func ReverseAry(s []int) {
	fmt.Printf("before Reverse:%v\n", s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("after Reverse:%v\n", s)
}

func CreateSliceByMake() {
	//create slice by make
	temp := make([]int, 5, 10)
	fmt.Printf("slice[len:%d\t/cap:%d\t] --> %v\n", len(temp), cap(temp), temp)

	temp1 := make([]int, 5)
	fmt.Printf("slice[len:%d\t/cap:%d\t] --> %v\n", len(temp1), cap(temp1), temp1)
	fmt.Printf("%v\n", temp1[3:])

	fmt.Printf("create empty slice\n")
	nilSlice := make([]string, 0)
	fmt.Printf("empty slice[len:%d/cap:%d/nil:%t] --> %v\n", len(nilSlice), cap(nilSlice), nilSlice == nil, nilSlice)
}

func ResizeArray(x []int, y ...int) []int {
	var z []int
	var zlen = len(x) + len(y)
	if zlen <= cap(x) { //+1之后的len<=cap的话，说明可以直接使用，否则需要重新构造一个更大的slice
		z = x[:zlen] //slice x的底层数组还有空间，所以直接slice到zlen的位置
	} else {
		//slice的底层数组没有空间，需要重新申请一个
		capz := zlen
		if zlen < (capz << 1) {
			capz = capz << 1
		}

		z = make([]int, zlen, capz) //重新构造一个切片
		copy(z, x)
	}
	//从len(x)开始，将slice y拷贝到新的slice z中
	copy(z[len(x):], y)
	return z
}

func AppendString(info string){
	var runes []rune
	for _,code := range info{
		runes = append(runes,code)
	}
	fmt.Printf("string --> char slice : %q\n",runes)
}

func TrimEmpty(info []string) []string{
	var i int
	for _,str := range info{
		if str != "" && str != " "{
			info[i] = str
			i++
		}
	}
	return info[:i]
}

func RemoveItemInSlice(s []int,r int) []int{
	if r >= len(s){
		fmt.Printf("remove index should be from %d to %d\n",0,len(s) - 1)
		return nil
	}

	copy(s[r:],s[r+1:])

	return s[:len(s) - 1]
}
