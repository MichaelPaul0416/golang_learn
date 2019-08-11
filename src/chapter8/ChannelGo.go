package chapter8

import (
	"net"
	"log"
	"io"
	"os"
	"fmt"
	"time"
)

func ClientToServer() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan string) //获取一个string类型的通道
	go func() {
		fmt.Println("goroutinue")
		io.Copy(os.Stdout, conn) //把
		fmt.Println("done")
		done <- "done"
	}()

	fmt.Printf("before MustCopy\n")
	MustCopy(conn, os.Stdin) //stdin的流写入到conn中
	fmt.Printf("after MustCopy\n")
	conn.Close()

	fmt.Printf("before receive data from channel")
	<-done

}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Printf("copy error:%v\n", err)
		log.Fatal(err)
	}
}

//使用三个角色，两个通道
func Calculator(b bool) {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 9; x++ {
			if x > 3 && b {
				//此时需要关闭通道，后续如果在关闭的通道上继续发送数据的话，就会报错
				fmt.Printf("close channel:%s\n", "naturals")
				close(naturals)
			}
			naturals <- x //将x发送到通道中
			time.Sleep(time.Second * 1)
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals //从通道中接收值
			//返回的ok代表通道上的数据全部被接受完毕，并且通道已经被关闭
			if !ok {
				fmt.Printf("data in channel has been all received and the channel closed...\n")
				break
			}
			squares <- x * x //将计算好的平方重新发送到新的通道中
		}
		close(squares)
	}()

	//printer 在主goroutine
	for x := range squares {
		fmt.Printf("%d\n", x)
	}
}

//out：单向通道，只能发送
func Counter(out chan<- int, v int) {
	//单向，发送
	for i := 0; i < v; i++ {
		//fmt.Printf("produce and send:%d\n", i)
		out <- i
	}
	fmt.Printf("all produced and close channel\n")
	close(out)//如果这里不关闭，那么从这个通道等待接收的goroutine将会被阻塞
}

func ReceiveAndCalculator(out chan<- int, in <-chan int) {
	for v := range in { //range 通道接受完毕并且通道关闭
		//fmt.Printf("receive item:%d\n", v)
		r := v * v
		//fmt.Printf("calculator and send:%d\n", r)
		out <- r
	}
	//如果Counter中不close的话，那么这一行代码永远不会被执行，当channel的数据被消费完毕之后，执行该函数的goroutine将会一直被阻塞，直到channel中有新数据进入
	fmt.Printf("close middle channel\n")
	close(out)
}

func FinalResult(in <-chan int) {
	for f := range in {
		fmt.Printf("%d\n", f)
	}
}

//带有缓冲区的通道的声明
func ProducerAndConsumer(l, num int) {
	ch := make(chan int, l) //创建一个容量为l的int类型的channel

	go func() {
		for i := 0; i < num; i++ {
			ch <- i
		}
		fmt.Printf("all send and will close channel...\n")
		close(ch)
	}()

	for r := range ch {
		fmt.Printf("%d\n",r)
		time.Sleep(time.Second * 1)
	}
}
