package chapter4

import (
	"encoding/json"
	"fmt"
	"time"
	"net/http"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title:"Casablanca",Year:1942,Color:false,Actors:[]string{"A","B"}},
	{Title:"复联4",Year:2019,Color:true,Actors:[]string{"小罗伯特唐尼","斯佳丽 约翰逊"}},
}

const IssuesURL = "https://api.github.com/repos/vmg/redcarpet/issues?state=closed"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string)(*IssuesSearchResult,error){
	//q := url.QueryEscape(strings.Join(terms," "))
	//fmt.Printf("url:%s\n",IssuesURL + "?q=" +q)
	resp,err := http.Get(IssuesURL)
	if err != nil{
		fmt.Printf("error http get issues:%v\n",err)
		return nil,err
	}

	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil,fmt.Errorf("search query failed:%s",resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err!=nil{
		resp.Body.Close()
		return nil,err
	}

	resp.Body.Close()
	return  &result,nil
}

func ObjectToJson(format bool){
	var b []byte
	if format{
		data,err := json.MarshalIndent(movies,""," ")
		if err != nil {
			fmt.Printf("error while transfer to json:%v\n",movies)
			return
		}
		b = data
		fmt.Printf("json:%s\n",data)
	}else {
		data, err := json.Marshal(movies)
		if err != nil {
			fmt.Printf("error while transfer to json:%v\n",movies)
			return
		}
		b = data
		fmt.Printf("json:%s\n",data)
	}
	var m []Movie
	e := json.Unmarshal(b,&m)
	if e != nil{
		fmt.Printf("json unmarshal failed:%v\n",e)
		return
	}

	fmt.Printf("unmarshal json:%v\n",m)

}
