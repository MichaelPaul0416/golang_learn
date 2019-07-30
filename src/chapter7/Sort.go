package chapter7

import (
	"sort"
	"fmt"
	"time"
	"text/tabwriter"
	"os"
	"bytes"
)

type StringSlice []string

//为StringSlice实现sort.Sort接口[Len，Less，Swap三个方法]
func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortSlice(s []string) {
	ss := StringSlice(s)
	sort.Sort(ss)
	fmt.Printf("sorted slice:%v\n", ss)
}

type Track struct {
	Title  string
	Artist string //作者
	Album  string //专辑
	Year   int
	Length time.Duration //时长
}

var tracks = []*Track{
	{"Go", "Delilah", "From The Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//将播放列表输出为一个表格
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}

	//格式化，并输出
	tw.Flush()
}

//定义能将Track按照Artist进行排序的封装对象
type byArtist []*Track

func (ba byArtist) Len() int {
	return len(ba)
}

func (ba byArtist) Less(i, j int) bool {
	return ba[i].Artist < ba[j].Artist
}

func (ba byArtist) Swap(i, j int) {
	ba[i].Artist, ba[j].Artist = ba[j].Artist, ba[i].Artist
}

func (t Track) String() string{
	var buf bytes.Buffer
	buf.WriteString("Track{")
	buf.WriteString("Title:"+t.Title)
	buf.WriteByte('\t')
	buf.WriteString("Artist:" +t.Artist)
	buf.WriteByte('\t')
	buf.WriteString("Album:" + t.Album)
	buf.WriteByte('\t')
	fmt.Fprintf(&buf,"Year:%d\t",t.Year)
	fmt.Fprintf(&buf,"Length:%s}",t.Length)

	return buf.String()
}

func SortTracksByArtist(){
	ba := byArtist(tracks)//构造原始的对象
	sort.Sort(ba)
	//按照Artist的字典升序排列
	printTracks(ba)

	//将其按照Artist的字典降序排列
	rba := byArtist(tracks)
	sort.Sort(sort.Reverse(rba))
	printTracks(rba)
	//for _,t := range ba {
	//	fmt.Printf("sorted by artist:%v\n", t.String())
	//}
}


//sort.Sort(x)的x不一定个都是slice，也可以是struct
type customSort struct {
	t []*Track
	less func(x,y *Track) bool
}

func (cs customSort) Len() int{
	return len(cs.t)
}

func (cs customSort) Less(i,j int) bool{
	return cs.less(cs.t[i],cs.t[j])//最终去执行成员变量函数less
}

func (cs customSort) Swap(i,j int){
	cs.t[i],cs.t[j] = cs.t[j],cs.t[i]
}

func SortWithMulti(){
	cm := customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title{
			return x.Title < y.Title
		}
		if x.Year != y.Year{
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}}

	sort.Sort(cm)
	printTracks(cm.t)
}