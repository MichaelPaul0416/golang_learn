package main

import (
	"image"
	"image/color"
	"fmt"
	"io/ioutil"
	"os"
	"math"
	"net/http"
	"html/template"
	"time"
	"strconv"
	"bytes"
	"image/jpeg"
	"encoding/base64"
	"sync"
	"image/draw"
)

// 马赛克图片的web服务

// 计算图片的平均颜色
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()                  //当前像素点的rgb值
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1) // 累计总和
		}
	}

	totalPixels := float64(bounds.Max.X * bounds.Max.Y) // 图片中总的像素量
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

// 把图片缩放至指定的宽度
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth

	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio,
		bounds.Max.X/ratio, bounds.Max.Y/ratio))

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

// 扫描瓷砖图片所在的目录创建一个瓷砖图片数据库
func tilesDB() map[string][3]float64 {
	fmt.Println("start population tiles db...")
	db := make(map[string][3]float64)

	files, _ := ioutil.ReadDir("~/devtool/tiles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)

		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in population TileDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("finished populating tiles db.")
	return db
}

// 寻找与目标图片相匹配的瓷砖图片
func (db *Db) nearest(target [3]float64) string {
	var filename string
	db.mutex.Lock()
	smallest := 1000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}

	delete(db.store, filename)
	db.mutex.Unlock()
	return filename
}

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}

var GlobalTilesDB map[string][3]float64

func closeTilesDB() Db {
	db := make(map[string][3]float64)
	for k, v := range GlobalTilesDB {
		db[k] = v
	}

	database := Db{&sync.Mutex{},db}
	return database
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	GlobalTilesDB = tilesDB()
	fmt.Println("Mosaic server started...")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

type Db struct {
	mutex *sync.Mutex
	store map[string][3] float64
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()

	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := closeTilesDB()

	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)

	c := combine(bounds, c1, c2, c3, c4)


	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":  <- c,
		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("result.html")
	t.Execute(w, images)
}

// 返回一个只能接受数据的单向通道
func cut(original image.Image, db *Db, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)
	sp := image.Point{0,0}
	go func() {
		newimage := image.NewNRGBA(image.Rect(x1,y1,x2,y2))
		for y:=y1;y<y2;y=y+tileSize{
			for x := x1;x<x2;x=x+tileSize{
				r,g,b,_ := original.At(x,y).RGBA()
				color := [3]float64{float64(r),float64(g),float64(b)}
				nearest := db.nearest(color)

				file,err := os.Open(nearest)
				if err == nil{
					img,_,err := image.Decode(file)
					if err == nil{
						t := resize(img,tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x,y,x+tileSize,y+tileSize)
						draw.Draw(newimage,tileBounds,tile,sp,draw.Src)
					}else{
						fmt.Println("error:",err)
					}
				}else{
					fmt.Println("error:",nearest)
				}
				file.Close()
			}
		}

		c <- newimage.SubImage(newimage.Rect)
	}()

	return c
}


// 返回一个接收数据的单向通道
func combine(r image.Rectangle,c1,c2,c3,c4 <- chan image.Image) <- chan string{
	c := make(chan string)
	go func() {
		var wg sync.WaitGroup
		img := image.NewNRGBA(r)
		copy := func(dst draw.Image,r image.Rectangle,src image.Image,sp image.Point) {
			draw.Draw(dst,r,src,sp,draw.Src)
			wg.Done()
		}

		wg.Add(4)
		var s1,s2,s3,s4 image.Image
		var ok1,ok2,ok3,ok4 bool
		for{
			select {
			case s1,ok1 = <- c1:
				go copy(img,s1.Bounds(),s1,image.Point{r.Min.X,r.Min.Y})
			case s2,ok2 = <- c2:
				go copy(img,s2.Bounds(),s2,image.Point{r.Max.X/2,r.Min.Y})
			case s3,ok3 = <- c3:
				go copy(img,s3.Bounds(),s3,image.Point{r.Min.X,r.Max.Y/2})
			case s4,ok4 = <- c4:
				go copy(img,s4.Bounds(),s4,image.Point{r.Max.X/2,r.Max.Y/2})
			}

			if ok1 && ok2 && ok3 && ok4{
				break
			}
		}

		wg.Wait()

		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2,img,nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()

	return c
}