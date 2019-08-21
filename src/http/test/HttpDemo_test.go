package test

import (
	"testing"
	//绝对路径导包，相当于$GOPATH/http
	"http"
)

func TestSumNumbers(t *testing.T) {
	if i := http.SumNumbers(1, 2); i != 3 {
		t.Error("error while call SumNumber")
	} else {
		t.Log("ok")
	}
}
