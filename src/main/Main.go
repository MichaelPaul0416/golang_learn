package main

import (
	"../chapter5"
	"os"
	"fmt"
	"net/http"
	"io"
	"../golang.org/x/net/html"
	"strings"
)

func main() {
	//p := new(int)
	//fmt.Printf("address:%v\n", p)
	//fmt.Printf("value:%v\n", *p)
	//*p = 123
	//fmt.Println("change value...")
	//fmt.Printf("address:%v\n", p)
	//fmt.Printf("value:%v\n", *p)
	//
	//fmt.Println( chapter2.CToF(100))
	//boilingF := chapter2.CToF(100)
	//fmt.Printf("%g\n",boilingF - chapter2.CToF(chapter2.FreezingC))
	////compiler failed
	////fmt.Printf("%g\n",BoilingC - CToF(100))
	//
	//c := chapter2.FToC(212)
	//fmt.Println(c.String())
	//fmt.Printf("%v\n",c)//调用String方法返回格式
	//fmt.Printf("%g\n",c)//不输出格式

	//---------------------------chapter3---------------------------
	//chapter3.ShowNumber()
	//
	//s := chapter3.BaseName("/home/george/IdeaProjects/JavaKnowledge/go/main/a.b.c.d.go")
	//fmt.Println(s)
	//s = chapter3.BaseNameWithLib("/home/george/IdeaProjects/JavaKnowledge/go/main/a.b.c.d.go")
	//fmt.Println(s)
	//s = chapter3.FormatNumber("12345689")
	//fmt.Println(s)
	//
	//chapter3.ConvertToString(123)
	//chapter3.ShowConstant()

	//---------------------------chapter4----------------------------
	//chapter4.DefArray()
	//var a = [3]int{1, 2, 3}
	//chapter4.ChangeArray(a, 2)
	//fmt.Printf("方法外数组：%v\n", a)
	//chapter4.ChangeArrayByPoint(&a, 2)
	//fmt.Printf("指针修改，方法外数组：%v\n", a)
	//
	//chapter4.ShowNumberMap()
	//chapter4.ShowSliceFromArray(3, 5)
	////数组字面量
	////ary := [...]int{1,2,3,4,5,6,7,8}
	//
	////slice字面量
	//ary := []int{1, 2, 3, 4, 5, 6, 7, 8}
	////slice 包含了指向数组元素的指针
	//chapter4.ReverseAry(ary[1:])
	//
	//fmt.Printf("now:%v\n", ary)
	////将ary中的所有元素都左移3位
	//chapter4.ReverseAry(ary[:3])
	//chapter4.ReverseAry(ary[3:])
	//chapter4.ReverseAry(ary)
	//
	////通过make创建slice
	//chapter4.CreateSliceByMake()
	//resize := chapter4.ResizeArray([]int{1,2,3,4},5)
	//fmt.Printf("resize slice:%v\tlen:%d\tcap:%d\n",resize,len(resize),cap(resize))
	//chapter4.AppendString("hello world,你好啊")
	//
	//str := []string{"1","","2"," ","3"}
	//fmt.Printf("before trim : %v\n",str)
	//str = chapter4.TrimEmpty(str)
	//fmt.Printf("[]string:%v\n",str)
	//
	//r := []int{1,2,3,4,5,6,7,8,9}
	//fmt.Printf("before remove:%v\n",r)
	//r = chapter4.RemoveItemInSlice(r,5)
	//fmt.Printf("after remove:%v\n",r)
	//
	//chapter4.InitMap()
	//m := make(map[string]int)
	//m["b"] = 2
	//m["a"] = 1
	//m["c"] = 3
	//chapter4.PrintMap(m)
	//
	////对key进行排序
	//k := chapter4.SortedKeys(m)
	////k --> []string
	//for _,key := range k{
	//	fmt.Printf("key-->%v/value-->%d\t",key,m[key])
	//}
	//fmt.Printf("\n")
	//
	////map -- nil
	//chapter4.NilMap()
	//
	////key exist
	//chapter4.ExistKey(m,"c")
	//chapter4.ExistKey(m,"d")
	//
	////map equals
	//n := make(map[string]int)
	////n["c"] = 3
	//n["b"] = 2
	//n["a"] = 1
	//fmt.Printf("map equals:%t\n",chapter4.EqualsMap(m,n))
	//
	//c := make(map[string]map[string]bool)
	//l1 := make(map[string]bool)
	//l1["h"] = true
	//l1["e"] = true
	//c["A"] = l1
	//chapter4.MapList("A","d",c)

	//var employee chapter4.Employee
	//employee.Address = "Hangzhou"
	//employee.ID = 1
	//employee.ManagerID = 0
	//employee.Name = "Bob"
	//employee.Position = "CTO"
	//employee.Salary = 20000
	//chapter4.Describe(&employee)
	//fmt.Printf("inner change:%v\n",employee)
	//
	//var tr = [9]int{8,4,5,7,1,9,11,10,12}
	//or := chapter4.Sort(tr[:])
	//chapter4.PrintTreeBefore(or)
	//fmt.Println()
	//
	//var sc = chapter4.Scale{1,2}
	//fmt.Printf("%v\n",chapter4.LargerScale(sc,3))
	//chapter4.LargerScaleByPoint(&sc,5)
	//fmt.Printf("larger by point:%v\n",sc)
	//
	//chapter4.ShowWheel()
	//
	//chapter4.ObjectToJson(true)

	//result,err := chapter4.SearchIssues([]string{"repo:golang/go","is:open","json","decoder"})
	//if err != nil{
	//	fmt.Printf("error:%v\n",err)
	//	return
	//}

	//fmt.Printf("%d issues:\n",result.TotalCount)
	//for _,item := range result.Items{
	//	fmt.Printf("#%-5d %9.9s %.55s\n",
	//		item.Number,item.User.Login,item.Title)
	//}

	//var str = [...]string{"a","b","c","d","e"}
	//fmt.Printf("str:%s\n",strings.Join(str[:],","))
	//
	//chapter4.GenerateFromTemplate()

	//-----------------------chapter5-----------------------
	chapter5.ShowConst()

	var s = [20]int32{}
	sl := s[20:]
	for i := 0; i < 10; i++ {
		sl = chapter5.ChangeStack(sl, int32(i))
	}

	//code := fetch("https://github.com")
	code := "<!DOCTYPE html><html lang=\"en\"><head></head><body class=\"logged-out env-production page-responsive f4\"><div class=\"position-relative js-header-wrapper \"><a href=\"#start-of-content\" tabindex=\"1\" class=\"px-2 py-4 bg-blue text-white show-on-focus js-skip-to-content\">Skip to content</a></body></html>";
	doc, err := html.Parse(strings.NewReader(code))
	//println(len(code))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range chapter5.Visit(nil, doc) {
		fmt.Println(link)
	}

	//chapter5.FetchAndParse("https://github.com/MichaelPaul0416?tab=repositories")

	//if line,err := chapter5.ReadUntilEOF();err != nil{
	//	fmt.Printf("error:%v\n",err)
	//}else{
	//	fmt.Printf("line:%s\n",line)
	//}

	//函数变量
	f := chapter5.Square
	fmt.Printf("square:%d\n", f(3))
	f = chapter5.Negative
	fmt.Printf("negative:%d\n", f(10))
	//不能进行如下赋值，会导致编译报错，因为f的类型是func (int) int 而chapter5.Product的类型却是func (int int) int
	//f = chapter5.Product

	r := chapter5.Composite(3, func(n int) int {
		return n * n
	})
	fmt.Printf("composite:%d\n", r)

	//匿名函数
	w := chapter5.WithoutNameFunc(2)
	fmt.Printf("func without name:%d\n", w(3))

	cur := chapter5.TopoSort(chapter5.Prereqs)
	fmt.Printf("cur:%v\n", cur)

	l1 := chapter5.TopoSort(chapter5.Prereqs)
	for i, course := range l1 {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

}

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http error :$%v\n", err)
		os.Exit(1)
	}

	writer, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "io error :%v\n", err)
		os.Exit(1)
	}

	return fmt.Sprintln(writer)
}
