package chapter8

import (
	"fmt"
	"time"
	"flag"
	"math/rand"
	"bytes"
	"sync"
	"os"
)

func SelfCycle() bool {
	flag := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		flag <- struct{}{}
	}()

	//如果在没有default的情况下,需要跳出select+for的话,需要使用标签,参考WatchDirectory.go
	for {
		//如果没有外面的for,那么select只会被执行一次,要么是case要么是default
		select {
		case <-flag:
			return true
		default:
			fmt.Printf("do not receive signal and sleep for a while\n")
			time.Sleep(1 * time.Second)
		}
	}

	return false
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		fmt.Printf("cancel action...\n")
		return true
	default: //如果上面的不符合,那么这里立即返回
		return false
	}
}

var vr = flag.Bool("vr", false, "show processing")

//定义一个全局的令牌桶,当获取到令牌了,才可以开始下载
var tokens = make(chan struct{}, 20)

type worker struct {
	name string
}

//implement String method
func (w worker) String() string {
	var buf bytes.Buffer
	buf.WriteString("worker:")
	buf.WriteString(w.name)
	return buf.String()
}

//模拟一个很耗时的操作
func DownloadBigFile() {
	flag.Parse() //放在最前面,不然下面的vr就解析不了
	percent := make(chan float64)
	var tick <-chan time.Time

	if *vr {
		fmt.Printf("show processing...\n")
		tick = time.Tick(1000 * time.Millisecond)
	}

	urls := flag.Args()
	if len(urls) == 0 {
		urls = []string{"https://www.baidu.com"}
	}

	var pro float64 = 0

	//开启一个goroutine,检测取消动作
	go func() {
		os.Stdin.Read(make([]byte,1))
		close(done)
	}()

	//计数器,goroutine的数目,当计数器=0时,代表所有goroutine计算完毕
	var lock sync.WaitGroup
	for i, url := range urls {
		//如果此时取消了,那么直接返回
		if cancelled(){
			return
		}

		//goroutine
		name := fmt.Sprintf("download-%d", i)
		w := worker{name: name}
		lock.Add(1)
		go w.batchDownload(url, percent, &lock) //分发给对象,在子goroutine中执行
	}

	go func() {
		lock.Wait()
		close(percent)
	}()

	//main goroutine负责打印进度,感知是否有取消操作,感知chan是否关闭
	receiveAndPrintProcessing(&pro, percent, tick)

	fmt.Printf("total download %.1f%%\n",pro)
}

func (w worker) batchDownload(url string, percent chan float64, lock *sync.WaitGroup) {
	//为了控制并发,避免生成太多的goroutine
	//0.计数器+1,在本方法的外部被调用
	//1.先获取token,然后再计算
	//2.计算完了之后归还token
	//3.计数器-1
	//而defer的执行顺序是从下往上的,所以计数器-1需要在最后面执行

	defer lock.Done()

	//获取token和子goroutine感知取消在一个select中,当取消了,麻烦返回,不获取token
	select {
	case tokens <- struct{}{}:
		fmt.Printf("[%s] get token and start download...\n",w)
	case <-done:
		return
	}

	defer func() {
		<-tokens
	}()

	fmt.Printf("[%s] download resource from:%s\n", w, url)
	r := rand.Intn(5000)
	t := time.Duration(r)
	fmt.Printf("[%s] cost %d milliseconds\n", w, r)

	//sleep to look like work time
	time.Sleep(t * time.Millisecond)

	cost := float64(r) / 100
	percent <- cost
	fmt.Printf("[%s] download %.1f%%\n", w, cost)

}

//data是[]string类型的通道,tick是从time.Time类型的chan中接受数据
func receiveAndPrintProcessing(t *float64, percent chan float64, tick <-chan time.Time) {
loop:
	for {
		select {
		case <-done: //新增一个感知取消的操作
			for lf := range percent {
				fmt.Printf("left unreceived:%.1f%%\n",lf)
			}
			fmt.Printf("call cancel \n")
			return
		case p, ok := <-percent:
			if !ok {
				break loop //通道关闭
			}
			*t += p //进度累加

		case <-tick:
			fmt.Printf("current percent:%.1f%%\n", *t)
		}
	}
}

//test worker string method
func Worker() {
	w := worker{"Paul"}
	fmt.Printf("%s\n", w)

	//defer的顺序从下往上执行
	defer func() {
		fmt.Printf("1\n")
	}()

	defer func() {
		fmt.Printf("2\n")
	}()
}
