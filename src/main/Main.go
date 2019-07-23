package main

import (
	"../chapter7"
	"net/http"
	"log"
)

func main(){
	db := chapter7.Database{"shoes":50,"sock":12}
	//log.Fatal(http.ListenAndServe("localhost:8080",db))//直接所有业务都写在一个Handler中

	//每个url对应一和handler处理
	//mux := http.NewServeMux()
	mux := new(http.ServeMux)//和上面一行的效果是一样的
	//http.HandlerFunc其实是一个适配器，它实现了接口Handler的方法，同时自己作为一和函数类型
	//其实自己也可以实现一个，并在其中添加日志处理
	//下面这行的效果和被注释的那一行效果是一样的，都是将一个方法包装为Handler类型
	mux.Handle("/price",chapter7.HandlerAdapter(db.Price))
	//mux.Handle("/price",http.HandlerFunc(db.Price))
	mux.Handle("/list",http.HandlerFunc(db.List))
	mux.Handle("/addOrUpdate",chapter7.HandlerAdapter(db.InsertOrUpdate))
	mux.Handle("/deleteItem",chapter7.HandlerAdapter(db.DeleteItem))
	//下面这个主要是熟悉error接口
	mux.Handle("/seeError",chapter7.HandlerAdapter(db.SeeError))//将传入的错误信息封装成一个error返回给调用端
	log.Fatal(http.ListenAndServe("localhost:8080",mux))

}
