package chapter7

import (
	"net/http"
	"fmt"
	"strconv"
	"errors"
)

type dollars float32

//String方法被默认调用的时机是fmt输出的时候，format里面格式化写的是%s
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f\n", d)
}

type Database map[string]dollars

func (db Database) List(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item") //获取get的参数
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item:%s\n", item)
		return
	}
	fmt.Fprintf(w, "price of item[%s]:%s\n", item, price)
}

func (db Database) InsertOrUpdate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	w.WriteHeader(http.StatusOK)
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "wrong price:%s\n", price)
		return
	}

	if _, ok := db[name]; !ok {
		fmt.Fprintf(w, "set %s for price:%s\n", name, price)
	} else {
		fmt.Fprintf(w, "update price[%s] for item[%s]\n", price, name)
	}
	db[name] = dollars(f)
}

func (db Database) DeleteItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		fmt.Fprintf(w, "item:%s not exist\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "delete item:%s successfully\n", item)
}

func (db Database) SeeError(w http.ResponseWriter, r *http.Request) {
	txt := r.URL.Query().Get("err")
	err := errors.New(txt)
	fmt.Fprintf(w,"error detail:%s\n",err)
}

//自定义适配器，适配Handler接口
type HandlerAdapter func(w http.ResponseWriter, r *http.Request)

func (h HandlerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("adapter handler call...\n")
	h(w, r)
}

//Http所有的业务url都写在一个方法里面
func (db Database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		//列出所有商品和单价
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item") //获取get的参数
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item:%s\n", item)
			return
		}
		fmt.Fprintf(w, "price of item[%s]:%s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such url:%s\n", r.URL.Path)
	}
}
