package chapter8

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"flag"
	"time"
	"log"
	"sync"
)

var verbose = flag.Bool("v", false, "show processing while counting per second")

//创建一个20个容量的缓冲区通道,其实就是i一个信号量,当都满了的时候,就代表下一个goroutine无法写入
var sema = make(chan struct{}, 20)

//直接输出统计的结果
func TotalFileSize() (i int64, err error) {
	flag.Parse() //开启命令行输入时,参数解析
	roots := preCheckDir()

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles ++
		nbytes += size
	}

	printDetails(nfiles, nbytes)
	return nfiles, nil
}

func preCheckDir() []string {
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	return roots
}

//一定周期内输出统计的进度
func TotalFileSizeByProcess() {
	flag.Parse()
	dirs := preCheckDir()
	fmt.Printf("look dirs:%v\n", dirs)

	fileSize := make(chan int64)
	var tick <-chan time.Time
	if *verbose {
		log.Println("show processing...")
		tick = time.Tick(1 * time.Second)
	}

	//使用goroutine异步去统计,main的goroutine不至于被阻塞
	go func() {
		for _, dir := range dirs {
			walkDir(dir, fileSize)
		}

		//因为每一个需要遍历的根目录相对来说都是同步的,所以for执行完了说明结果都完事了
		close(fileSize)
	}()

	var files, bytes int64
	processing(fileSize, &files, &bytes, tick)
	printDetails(files, bytes)
}

func processing(fileSize chan int64, files *int64, bytes *int64, tick <-chan time.Time){
loop:
	for {
		select { //注意,select不是一个循环,而是一次过程,只不过可能是会阻塞的,所以一般外面都是需要for配合
		case size, ok := <-fileSize:
			if !ok { //判断fileSizes通道是否已经关闭
				break loop //声明退出的是loop后面的for循环
				//如果注释掉loop:并且使用break,后面不跟loop的话,那么只是退出当前的select的逻辑,但是外面还有一层for,所以相当于又一次进入了select
			}
			*files ++
			*bytes += size
		case <-tick:
			printDetails(*files, *bytes)

		}
	}
}

//使用goroutine并发统计目录大小
func ConcurrencyDirSize() {
	flag.Parse()
	dirs := preCheckDir()

	//命令行加入了-v参数的话,就开启定时统计的功能
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(1 * time.Second)
	}

	fileSize := make(chan int64)
	var n sync.WaitGroup //计数监视器

	for _, dir := range dirs {
		//开启goroutine之前,需要先计数器+1
		n.Add(1)
		//每一层都开启一个goroutine
		go walkDirConcurrence(dir, &n, fileSize)
	}

	//开启一个新的goroutine负责关闭通道,条件是计数器为0
	go func() {
		n.Wait()
		close(fileSize)
	}()

	var files, bytes int64
	processing(fileSize, &files, &bytes, tick)

	printDetails(files, bytes)
}

func printDetails(nfiles int64, nbytes int64) (int, error) {
	return fmt.Printf("%d files / %.1f byte / %.1f Mb \n", nfiles, float64(nbytes), float64(nbytes)/1e6)
}

func walkDirConcurrence(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done() //递归调用的时候,不论是跟文件夹还是子文件夹,都会执行
	for _, entry := range dirListBySema(dir) {
		if entry.IsDir() {
			n.Add(1)                                   //当前是一个文件夹,所以计数器需要再+1
			subdir := filepath.Join(dir, entry.Name()) //生成子文件夹的名字
			go walkDirConcurrence(subdir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirent(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//并发统计的时候,这个方法是每个goroutine都会进入的(dir只要是文件夹),而计算每一个当前文件夹的大小,都是由一个goroutine计算,所以在这里加入获取信号量
func dirListBySema(dir string) []os.FileInfo{
	sema <- struct{}{}//这个方法是每个

	defer func() {
		<-sema
	}()
	return dirent(dir)
}

func dirent(dir string) []os.FileInfo {

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}

	return entries
}
