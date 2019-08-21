package test

import (
	"testing"
	"time"
	"web"
)

func TestStartServer(t *testing.T){
	web.StartServer()
	for {
		time.Sleep(10 * time.Hour)
	}
}
