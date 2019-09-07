package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post1 struct {
	//`xml:"post"`:结构标签，使用这些标签决定如何对结构以及XML元素进行映射
	//结构标签的形式：`key:"value"`
	XMLName xml.Name `xml:"post"`
	//将Post节点的id属性对应的值设置为结构体Post的Id值
	Id      string   `xml:"id,attr"`
	//建立xml的content节点到Post结构体中Content字段的映射
	Content string   `xml:"content"`
	Author  Author1   `xml:"author"`
	//post标签内部的xml报文，将Post节点的xml报文存储到这个字段里面
	Xml     string   `xml:",innerxml"`

	//使用类似java的xpath访问,将post节点下路径为comments/comment的节点分析组装到这个属性里面
	Comments []Comment1 `xml:"comments>comment"`
}

type Author1 struct {
	Id   string `xml:"id,attr"`
	//通过设置`xml:",chardata"`，将Author节点的标签值设置为Author结构体的Name属性值
	Name string `xml:",chardata"`
}

type Comment1 struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author1 `xml:"author"`
}

func Encoder(post1 *Post1){
	xmlFile,err := os.Create("XmlEncoder.xml")
	if err != nil{
		fmt.Printf("error creating xml:%v\n",err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("","\t")
	err = encoder.Encode(post1)
	if err != nil{
		fmt.Printf("error encoding xml:%v\n",err)
		return
	}
}

func main(){
	post := Post1{
		Id:"1",
		Content:"Hello world!",
		Author:Author1{
			Id:"a-1",
			Name:"Jane",
		},
	}

	//生成的是一行xml,没有经过格式化的
	//output,err := xml.Marshal(&post)

	//生成的xml是按照xml的标准形式,经过格式化之后的
	output,err := xml.MarshalIndent(&post,"","\t")
	if err != nil{
		fmt.Printf("marshal xml error:%v\n",err)
		return
	}

	//生成的xml没有标准的xml文件头
	//err = ioutil.WriteFile("marshalXml.xml",output,0644)

	//添加标准的xml文件头
	err = ioutil.WriteFile("marshalXml.xml",[]byte(xml.Header + string(output)),0644)
	if err != nil{
		fmt.Printf("write data to xml failed:%v\n",err)
	}

	fmt.Printf("------------------------------------\n")
	Encoder(&post)
}