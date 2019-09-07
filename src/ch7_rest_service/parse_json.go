package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
)

type PostJson struct {
	Id      int   `json:"id"`
	Content string   `json:"content"`
	Author  AuthorJson   `json:"author"`
	Comments []CommentJson `json:"comments"`
}

type AuthorJson struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type CommentJson struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func decode(){
	jsonFile,err := os.Open("src/ch7_rest_service/post.json")
	if err != nil{
		fmt.Printf("open json file failed:%v\n",err)
		return
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for{
		var post PostJson
		err := decoder.Decode(&post)
		if err == io.EOF{
			break
		}

		if err != nil{
			fmt.Printf("decode json file error:%v\n",err)
			return
		}

		fmt.Printf("json:%v\n",post)
	}
}

/**
关于使用Unmarshal还是deocder,主要看json数据的来源是什么,如果是io.Reader流,那么用decoder,如果是字符串或者内存的某个地方,使用Unmarshal更好
 */
func main(){
	jsonFile,err := os.Open("src/ch7_rest_service/post.json")
	if err != nil{
		fmt.Printf("error open json file:%v\n",err)
		return
	}

	defer jsonFile.Close()
	jsonData,err := ioutil.ReadAll(jsonFile)
	if err != nil{
		fmt.Printf("error reading json data:%v\n",jsonData)
		return
	}

	var postJson PostJson
	json.Unmarshal(jsonData,&postJson)
	fmt.Printf("json:%v\n",postJson)

	fmt.Printf("---------------------------------\n")
	decode()
}
