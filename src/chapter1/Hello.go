package chapter1

import (
	"os"
	"fmt"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep+os.Args[i]
		sep = " "
	}
	fmt.Printf("len(os.Args)=%d\n",len(os.Args))
	fmt.Println(s)
	fmt.Println("os.Args[0]=" + os.Args[0])
	fmt.Println("-----使用_丢弃数组下标-----")

	echo,tmp := ""," "
	for _,value := range os.Args[1:]{
		echo += value + tmp
		tmp = " "
	}
	fmt.Println(echo)
}
