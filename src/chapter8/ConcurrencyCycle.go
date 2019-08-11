package chapter8

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

func WaitUtilAllDone(num int, e bool) error {
	//创建一个带有缓冲区的error类型的chan
	type item struct {
		n int
		e error
	}

	ch := make(chan byte)
	//下面关于 er chan error的都注释掉，不带缓冲区的chan相当于一个同步的chan
	//有多个error写入到chan中，那么chann其实只会接收一个【下方的被注释的if代码开始（if err := <-er; err != nil）】
	//var er chan error
	var err_item chan item

	if e {
		//er = make(chan error)
		err_item = make(chan item, num) //创建一个带有缓冲区的channel
	}

	for i := 0; i < num; i++ {
		//使用m作为形参接收，实参为后面的(i)，也就是此次循环的i变量的副本，保证在执行func的具体代码时，获取的是那次循环的i的值，而不是并发更新过后i的值
		//如果这里m去掉，并且后面的（i）也去掉，那么func中输出值基本都是3
		go func(m int) {
			fmt.Printf("show number:%d\n", m)
			//测试用
			if e {
				//手动抛出error
				fe := fmt.Errorf("wrong number:%d\n", m)
				//er <- fe

				temp := item{m, fe}
				err_item <- temp //这里用带缓冲区的chann来接收消息
			}
			time.Sleep(time.Second * 1)
			ch <- byte(1) //每个goroutine完成之后，向channel发送一个byte，表示搞完了
		}(i) //这里的这个i其实就是此次循环中的变量i，将其作为参数，传递给goroutine
	}

	if e {
		for k := 0; k < num; k++ {
			//if err := <-er; err != nil { //从专门接收error的channel中接收错误
				//不能这样写，这样的话只会从子goroutine中接收一个error，如果后续还有其他error的话，那么其他子goroutine利用er写入error的时候，会永久阻塞
				//因为er是不带缓冲区的channel
				//return err
				//这里直接用输出代替
				//fmt.Printf("error:%s\n", err)

				//直接输出的话，可能结果只会打印一行，因为一行error接收到之后，执行下面的if，下面带缓冲区的在接收到error之后，直接return了，所以就不会进行for的循环了
			//}

			if eh := <-err_item; eh.e != nil {
				//由于err_item是带有缓冲区的channel，所以即便这里直接返回，由于err_item有缓冲区，所以后续其他的goroutine向err_item发送错误消息时，也不会阻塞
				return eh.e
			}
		}
	} else {
		<-ch //channel阻塞接收byte，当所有的byte都接收完毕之后，返回
	}

	fmt.Printf("all number done...\n")
	return nil
}

//当不知道迭代的次数的时候，下面的代码结构可以通用
func CycleWhileUnConfirmTimes(times int) {
	var sequence []int
	sizes := make(chan int)  //接收返回的数据
	var watch sync.WaitGroup //类似Java的CountDownWatch
	for i := 0; i < times; i++ {
		watch.Add(1)
		go func(m int) {
			defer watch.Done()//保证即使发生错误也能被执行到

			fmt.Printf("number:%d\n",m)
			t := rand.Intn(1000)
			time.Sleep(time.Duration(t) * time.Microsecond)//随机休眠一段时间
			sizes <- m
		}(i)
	}

	//closer
	go func() {
		watch.Wait()
		close(sizes)
	}()

	for s := range sizes{
		sequence = append(sequence,s)
	}

	fmt.Printf("receive order:%v\n",sequence)

}
