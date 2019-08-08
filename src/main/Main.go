package main

import (
	"../chapter8"
	"fmt"
)

func main() {
	chapter8.Rocket()
	fmt.Println("------------------------")
	chapter8.RocketWithSelect()
	fmt.Println("------------------------")
	chapter8.ShowNumber(10,true)
	fmt.Println("------------------------")
	chapter8.MultiCasesSelector(10)
	fmt.Println("------------------------")
	chapter8.SendMessageWithTick(3)
	fmt.Println("------------------------")
	chapter8.UnBlockingSelector()
}
