package chapter4

import (
	"fmt"
	"sort"
)

func InitMap(){
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Printf("map:%v\n",m)

	fmt.Printf("delete key=two\n")
	delete(m,"one")
	fmt.Printf("map:%v\n",m)
}

func PrintMap(m map[string]int){
	for k,v := range m{
		fmt.Printf("key=%v\tvalue=%d\n",k,v)
	}
}

//入参为map，出参为排序后的key[]
func SortedKeys(m map[string]int) []string{
	var keys []string
	//只需要获取map的key，所以忽略map的value，for中就不体现
	for name := range m{
		keys = append(keys,name)
	}
	sort.Strings(keys)

	return keys
}

func NilMap(){
	var m map[string]int
	fmt.Printf("empty map:%t\n",m == nil)
	//nil的map不能设置k/v，设置之前map必须初始化
}

func ExistKey(m map[string]int,k string){
	//第二个值是bool，报告该元素是否存在
	v,ok := m[k]
	if !ok{
		fmt.Printf("key:%s not exist...\n",k)
	}else {
		fmt.Printf("key:%s/value:%d\n",k,v)
	}
}

func EqualsMap(m,n map[string]int) bool{
	if len(m) != len(n){
		return false
	}
	for k,mv := range m{
		if nv,ok := n[k]; !ok || mv != nv{
			return false
		}
	}
	return true
}

func MapList(from,to string,ml map[string]map[string]bool){
	if _,ok := ml[from];!ok{
		fmt.Printf("value for from:%s is empty\n",from)
		a := make(map[string]bool)
		ml[from] = a
	}
	fmt.Printf("complete map:%v\n",ml)
	a := ml[from]
	if !a[to]{
		fmt.Printf("code:%s not exist where key=%s\n",to,from)
	}else {
		fmt.Printf("code:%s in list where key=%s\n",to,from)
	}


}
