package main

import (
	"../chapter9"
	"net/http"
	"io/ioutil"
	"time"
	"log"
	"fmt"
	"sync"
)

func main() {
	//m := chapter9.New(httpGetBody) //绑定函数到memo,赋值为对应的成员变量
	//commonSerialTest(m)
	//multiGet:false的时候,使用粒度较大的锁,锁住Get整个方法,true的时候使用两把锁,分别锁住查询memo.cache和更新memo.cache两个过程
	//selfCycleConcurrency(m,true)

	//使用新版的memo2对象
	//m2 := chapter9.New2(httpGetBody)
	//selfCycleWithChannel(m2)

	//使用monitor goroutine进行处理的版本
	m3 := chapter9.New3(httpGetBody)
	selfCycleWithMonitorChannel(m3)
}

func selfCycleWithMonitorChannel(m3 *chapter9.Memo3) {
	var w sync.WaitGroup
	for key := range incomingUrls() {
		for i := 0; i < 3; i++ {
			w.Add(1)
			go func(m string) {
				defer w.Done()
				start := time.Now()
				value, err := m3.GetFromMonitorGoroutine(m)
				if err != nil {
					log.Print(err)
				}
				fmt.Printf("url[%s]\ttime[%s]\tbytes[%d]\n", m, time.Since(start), len(value.([]byte)))
			}(key)
		}
	}
	w.Wait()
}

func selfCycleWithChannel(m *chapter9.Memo2) {
	var w sync.WaitGroup
	for key := range incomingUrls() {
		for i := 0; i < 3; i++ {
			w.Add(1)
			go func(k string) {
				defer w.Done()
				start := time.Now()
				value, err := m.GetFromChannel(k)
				if err != nil {
					log.Print(err)
				}
				fmt.Printf("url[%s]\ttime[%s]\tbytes[%d]\n", k, time.Since(start), len(value.([]byte)))
			}(key)
		}
	}
	w.Wait()
}

/**
现在并发的执行,但是此时原来的缓存几乎没用,每个url因为缓存中get不到,而重新建立一个连接去访问远程的http服务器
 */
func selfCycleConcurrency(m *chapter9.Memo, multiGet bool) {
	var w sync.WaitGroup
	for key := range incomingUrls() {
		for i := 0; i < 3; i++ {
			w.Add(1)
			go func(k string) {
				defer w.Done()
				start := time.Now()
				//当Memo加上锁之后,同一个url多次访问输出的时间是接近的,因为当一个goroutine获取lock之后,其他goroutine是在等的,但是他们却已经执行了
				//time.Now方法,也就是已经开始计时了,当此时一个goroutine返回之后,其他对于同一个url访问的goroutine几乎是立刻返回,所以时间上和第一个类似
				var value interface{}
				var err error
				if multiGet {
					value, err = m.GetByTwoStep(k)
				} else {
					value, err = m.Get(k)
				}
				if err != nil {
					log.Print(err)
				}
				fmt.Printf("url[%s]\ttime[%s]\tbytes[%d]\n", k, time.Since(start), len(value.([]byte)))
			}(key)
		}
	}
	w.Wait()
}

func commonSerialTest(m *chapter9.Memo) {
	for l := range incomingUrls() {
		selfCycle(m, l)
	}
}

func selfCycle(m *chapter9.Memo, l string) {
	for i := 0; i < 3; i++ {
		start := time.Now()
		value, err := m.Get(l)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("url[%s]\ttime[%s]\tbytes[%d]\n", l, time.Since(start), len(value.([]byte)))
		//使用类型推断,将value强转为[]byte
	}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingUrls() (map[string]bool) {
	list := map[string]bool{
		"https://www.baidu.com": true,
		"https://www.qq.com":    true,
		"https://study.163.com": true,
	}
	return list

}
