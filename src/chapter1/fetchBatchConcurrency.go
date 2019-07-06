package chapter1

import (
	"os"
	"fmt"
	"bufio"
	"time"
	"net/http"
	"io"
	"io/ioutil"
)

/**
 *并发获取
 */

func main() {
	if len(os.Args) < 2{
		fmt.Fprintf(os.Stderr,"param must contains url file location")
		os.Exit(1)
	}
	urls := os.Args[1]

	channel := make(chan string) //创建一个channel，接受string
	f, err := os.Open(urls)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error open file:%s\n", err)
	}

	start := time.Now()
	input := bufio.NewScanner(f)
	site := 0
	for input.Scan() {
		url := input.Text()
		site++
		go fetch(url, channel) //go关键字，启动一个协程处理，处理结果都使用channel这个通道返回
	}

	for i := 0; i < site; i++ {
		fmt.Printf("receive:%s\n",<- channel)
	}

	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", secs, nbytes, url)
}
