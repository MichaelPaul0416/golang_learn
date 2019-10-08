package main

import (
	"os"
	"syscall"
	"fmt"
	"os/signal"
	"os/exec"
	"errors"
	"bytes"
	"io"
	"strconv"
	"strings"
	"runtime/debug"
	"sync"
	"time"
)

func main() {
	//signalDemo()

	go func() {
		time.Sleep(5 * time.Second)
		sendSignal()
	}()

	handleSignal()
}

func showPid(){
	pids, err := getPid()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Printf("pid:%v\n", pids)
}

func signalDemo() {
	signRecv := make(chan os.Signal, 1)
	// 希望自行处理的信号
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	// 如果第二个参数为空的话，那么就自行处理所有的信号,kill/stop除外
	// 操作系统向当前进程发送指定信号时发出通知，将当前进程指定的信号放入通道中
	// 这样该函数的调用方就可以从signal接收通道中按顺序获取操作系统发来的信号并进行相对应的处理
	signal.Notify(signRecv, sigs...)
	// 从通道中接受数据
	for sig := range signRecv {
		fmt.Printf("received a signal:%s\n", sig)
	}
}
func handleSignal()  {
	signalRecv1 := make(chan os.Signal,1)
	sigs1 := []os.Signal{syscall.SIGINT,syscall.SIGQUIT}
	fmt.Printf("set notifaction for %s... [sigRecv1]\n",sigs1)
	signal.Notify(signalRecv1,sigs1...)

	signalRecv2 := make(chan os.Signal,1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("ser notification for %s...[sigRecv2]\n",sigs2)
	signal.Notify(signalRecv2,sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range signalRecv1{
			fmt.Printf("received a signal from sigRecv1:%s\n",sig)
		}
		fmt.Printf("end .[sigRecv1]\n")
		wg.Done()
	}()

	go func() {
		for sig := range signalRecv2{
			fmt.Printf("received a signal from sigRecv2:%s\n",sig)
		}
		fmt.Printf("end .[sigRecv2]\n")
		wg.Done()
	}()

	fmt.Println("wait for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Printf("stop notification...")
	signal.Stop(signalRecv1)
	close(signalRecv1)
	fmt.Printf("done .[sigRecv1]\n")
	wg.Wait()
}

func sendSignal(){
	defer func(){
		if err := recover();err != nil{
			fmt.Printf("fatal error:%s\n",err)
			debug.PrintStack()
		}
	}()

	pids,err := getPid()
	if err != nil{
		fmt.Printf("pid parsing error:%s\n",err)
		return
	}

	fmt.Printf("pid :%v\n",pids)
	for _,pid := range pids{
		// 返回pid对应的句柄
		proc,err := os.FindProcess(pid)
		if err != nil{
			fmt.Printf("process finding error:%s\n",err)
			return
		}

		sig := syscall.SIGQUIT
		fmt.Printf("send signal '%s' to process (pid=%d)...\n",sig,pid)
		// 向对应的进程发送信号
		err = proc.Signal(sig)
		if err != nil{
			fmt.Printf("signal sending error:%s\n",err)
			/**
			注释掉下面这个return的话,只要pid对应的进程找到,那么就一定会向对应的进程发送信号,而对应的监听通道上收到这个信号之后,就会做对应的处理
			在这的话监听signalRecv2这个chan的goroutine就会收到信号,然后打印日志
			 */
			//return
		}
	}
}

func getPid() ([]int, error) {
	// 定义一个cmd的切片
	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "go run"),
		exec.Command("awk", "{print $2}"),
	}

	output, err := runCmds(cmds)
	if err != nil {
		fmt.Printf("Command Execution Error:%s\n", err)
		return nil, err
	}

	var pids []int
	for _, str := range output {
		pid, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, err
}

func runCmds(cmds []*exec.Cmd) ([]string, error) {
	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("The cmd slice is invalid!")
	}

	first := true
	var output []byte
	var err error
	for _, cmd := range cmds {
		fmt.Printf("run command:%v\n", getCmdPlain(cmd))
		if !first {
			var stdinBuf bytes.Buffer
			// 前一步的输出作为这一步的输入
			stdinBuf.Write(output)
			cmd.Stdin = &stdinBuf
		}
		var stdoutBuf bytes.Buffer
		cmd.Stdout = &stdoutBuf
		// do cmd
		if err = cmd.Start(); err != nil {
			return nil, getError(err, cmd)
		}
		// wait for result
		if err = cmd.Wait(); err != nil {
			return nil, getError(err, cmd)
		}

		output = stdoutBuf.Bytes()

		if first {
			first = false
		}
	}

	var lines []string
	var outputBuf bytes.Buffer
	// 将最终的结果输入到缓冲区中
	outputBuf.Write(output)

	for {
		line, err := outputBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, getError(err, nil)
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func getCmdPlain(cmd *exec.Cmd) string {
	var buf bytes.Buffer
	buf.WriteString(cmd.Path)
	// 获取对应的参数
	for _, arg := range cmd.Args[1:] {
		buf.WriteRune(' ')
		buf.WriteString(arg)
	}
	return buf.String()
}

func getError(err error, cmd *exec.Cmd, extraInfo ...string) error {
	var errMsg string
	if cmd != nil {
		errMsg = fmt.Sprintf("%s [%s %v]", err, (*cmd).Path, (*cmd).Args)
	} else {
		errMsg = fmt.Sprintf("%s", err)
	}

	if len(extraInfo) > 0 {
		errMsg = fmt.Sprintf("%s (%v)", errMsg, extraInfo)
	}
	return errors.New(errMsg)
}
