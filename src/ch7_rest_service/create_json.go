package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PostJson1 struct {
	Id      int   `json:"id"`
	Content string   `json:"content"`
	Author  AuthorJson1   `json:"author"`
	Comments []CommentJson1 `json:"comments"`
}

type AuthorJson1 struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

type CommentJson1 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func encoder(pj *PostJson1){
	jsonfile,err := os.Create("post_encoder.json")
	if err != nil{
		fmt.Printf("error creating json file:%v\n",err)
		return
	}

	encoder := json.NewEncoder(jsonfile)
	//格式化
	encoder.SetIndent("","\t")
	err = encoder.Encode(pj)
	if err != nil{
		fmt.Printf("error encoding json to file:%v\n",err)
	}
}

func main(){
	postjson1 := PostJson1{
		Id:1,
		Content:"hello world",
		Author:AuthorJson1{
			Id:2,
			Name:"jane",
		},

		Comments:[]CommentJson1{
			CommentJson1{
				Id:3,
				Content:"have a good day",
				Author:"long",
			},
			CommentJson1{
				Id:4,
				Content:"how are you",
				Author:"netty",
			},
		},
	}

	output,err := json.MarshalIndent(&postjson1,"","\t")
	if err != nil{
		fmt.Printf("error marshalling to json:%v\n",err)
		return
	}

	err = ioutil.WriteFile("postJson.json",output,0644)
	if err != nil{
		fmt.Printf("error writing json to file:%v\n",err)
	}

	fmt.Printf("----------------------------------\n")
	encoder(&postjson1)
}