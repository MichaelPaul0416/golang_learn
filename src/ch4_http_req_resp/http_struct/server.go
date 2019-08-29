package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func showHeaders(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for h := range header{
		fmt.Fprintf(w, "%s:%s\n",h,header.Get(h))
	}
}

func body(w http.ResponseWriter,r *http.Request){
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	info := string(body)
	fmt.Printf("data:%s\n",info)
	fmt.Fprintf(w,info)
}

//http.Request.ParseForm --> [http.Request.Form|http.Request.FormValue|http.Request.PostForm|http.Request.PostFormValue]
func processForm(w http.ResponseWriter,r *http.Request){
	if r.ParseForm() != nil{
		fmt.Fprintf(w,"parse form failed\n")
		return
	}

	//这里将返回key-value，key可以同时存在url的键值对和form表单的字段中，也就是说如果一个key暨存在url中又存在form表单中，那么value将是同时包含对应值的slice
	urlValues := r.Form
	//r.Form返回的是key对应的所有值一个slice，r.FormValue(key)返回的是key对应的第一个值
	for k,v := range urlValues{
		fmt.Fprintf(w,"key:%s/value:%s\tsingleValue:%s\n",k,v,r.FormValue(k))
	}

	fmt.Fprintf(w,"--------------------输出结果只包含表单中的键值对--------------------\n")
	//PostForm只支持application/x-www-form-urlencoded类型的form表单
	formParams := r.PostForm
	//r.PostForm和r.PostFormValue(key)的效果和上面提到的r.Form/r.FormValue(key)是一样的
	for k,v := range formParams{
		fmt.Fprintf(w,"key:%s/value:%s\n",k,v)
	}
}

func processFormMulti(w http.ResponseWriter,r *http.Request){
	//先解析
	r.ParseMultipartForm(1024)

	//如果form表单的类型是multipart/form-data，那么想要获取表单里面的键值对，那么就需要使用MultipartForm对象
	fmt.Fprintf(w,"--------------------表单类型为multipart/form-data的时候，输出表单的键值对--------------------\n")
	multiForms := r.MultipartForm.Value
	for k,v := range multiForms{
		fmt.Fprintf(w,"key:%s/value:%s\n",k,v)
	}
}


func fileUpload(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(1024)
	//下面这两行代码也可以直接使用r.FormFile("name")代替
	//返回名字是upload的第一个附件
	//fileHeader := r.MultipartForm.File["upload"][0]
	//获取输入流
	//file,err := fileHeader.Open()

	file,_,err := r.FormFile("upload")

	if err != nil{
		fmt.Fprintf(w,"receive file error:%s\n",err)
		return
	}

	data,err := ioutil.ReadAll(file)
	if err == nil{
		fmt.Fprintln(w,"file content")
		fmt.Fprintln(w,string(data))
	}else{
		fmt.Fprintf(w,"read data error:%v\n",err)
	}
}
func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/headers/",showHeaders)
	http.HandleFunc("/body/",body)
	http.HandleFunc("/form/url",processForm)
	http.HandleFunc("/form/multi",processFormMulti)
	http.HandleFunc("/fileUpload",fileUpload)
	server.ListenAndServe()
}
