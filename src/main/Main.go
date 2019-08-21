package main

import (
	"web"
	"time"
)

func main()  {
	web.StartServer()
	for {
		time.Sleep(1 * time.Hour)
	}
}
