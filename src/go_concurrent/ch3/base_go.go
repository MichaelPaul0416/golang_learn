package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Printf("pid:%d\n",os.Getpid())
	fmt.Printf("parent pid:%d\n",os.Getppid())
}
