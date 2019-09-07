package main

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
	"bytes"
	"io"
)

type Post struct {
	//`xml:"post"`:结构标签，使用这些标签决定如何对结构以及XML元素进行映射
	//结构标签的形式：`key:"value"`
	XMLName xml.Name `xml:"post"`
	//将Post节点的id属性对应的值设置为结构体Post的Id值
	Id      string   `xml:"id,attr"`
	//建立xml的content节点到Post结构体中Content字段的映射
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	//post标签内部的xml报文，将Post节点的xml报文存储到这个字段里面
	Xml     string   `xml:",innerxml"`

	//使用类似java的xpath访问,将post节点下路径为comments/comment的节点分析组装到这个属性里面
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	//通过设置`xml:",chardata"`，将Author节点的标签值设置为Author结构体的Name属性值
	Name string `xml:",chardata"`
}

type Comment struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

func (author *Author) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Author{id:%s,Name:%s}\n", author.Id, author.Name)
	return buf.String()
}

func (post *Post) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf,"Post{XMLName:%s,Id:%s,Content:%s,Author:%s,Xml:%s}\n",
		post.XMLName,post.Id,post.Content,post.Author,post.Xml)
	return buf.String()
}

//当需要解析的xml是流式或者比较大的时候,使用用decoder进行逐步解析
func DecoderXml(){
	xmlFile,err := os.Open("src/ch7_rest_service/post.xml")
	if err != nil{
		panic(err)
	}

	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)
	for{
		t,err := decoder.Token()
		if err == io.EOF{
			fmt.Printf("read the end\n")
			return
		}

		if err != nil{
			fmt.Printf("error decoding xml into token :%v\n" , err)
			return
		}

		switch se:= t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment"{
				var comment Comment
				decoder.DecodeElement(&comment,&se)
				fmt.Printf("comment:%v\n",comment)
			}

			if se.Name.Local == "content"{
				var content string
				decoder.DecodeElement(&content,&se)
				fmt.Printf("content:%s\n",content)
			}

		}
	}
}

func main() {
	xmlFile, err := os.Open("src/ch7_rest_service/post.xml")
	if err != nil {
		fmt.Printf("open xml failed:%v\n", err)
		return
	}

	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Printf("read file data failed:%v\n", err)
		return
	}

	var p Post
	xml.Unmarshal(xmlData, &p)
	//fmt.Printf("xml data:%s\n", &p)
	fmt.Println(p)

	fmt.Printf("-------------------------------------------\n")
	DecoderXml()
}
