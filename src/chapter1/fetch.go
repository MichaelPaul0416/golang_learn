package chapter1

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

/**
 * 从制定的url对应的网页中获取数据
 */

func main()  {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr,"params must be <protocol url1 url2...>")
		os.Exit(1)
	}

	flag := os.Args[1]
	for _,url := range os.Args[2:]{
		if !strings.HasPrefix(url,"http://") || !strings.HasPrefix(url,"https://"){
			if flag == "1" {
				url = "http://" + url
			}else {
				url = "https://" + url
			}
		}
		resp,err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch:%v and http status:%s\n",err,resp.Status)
			os.Exit(1)
		}

		//b,err := ioutil.ReadAll(resp.Body)//需要下载整个数据流到缓冲区b
		writer,err := io.Copy(os.Stdout,resp.Body)//直接进行流拷贝，从src：resp.Body，拷贝到dest：os.Stdout
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch:reading %v and http status:%v\n",err,resp.Status)
			os.Exit(1)
		}

		//fmt.Printf("%s\n",b)//对应ioutil.ReadAll
		println(writer)
	}
}