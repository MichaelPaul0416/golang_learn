//找出重复行
package chapter1

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	//生成一个map，key为string，value为int，能做key的要求是这个类型可以进行==计算，有点类似java里面的map，key必须重写过hashcode
	/**
	 * 从标准输入中找出重复行
	 */
	//standardIO()

	/**
	 * 从文件中找出重复行
	 */
	 readFromFile()

}

func readFromFile()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("no file input...")
		return
	}

	for _, args := range files { //遍历每一个传入的文件路径
		f, err := os.Open(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	printRepeatLine(counts, 2)
}

func standardIO() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		in := input.Text()
		if in == "done" {
			break
		}
		counts[input.Text()] ++
	}
	printRepeatLine(counts, 2)
}

func printRepeatLine(counts map[string]int, times int) {
	for line, number := range counts {
		if number >= times {
			fmt.Printf("repeat line:%s-->%d\n", line, number)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
	}
}
