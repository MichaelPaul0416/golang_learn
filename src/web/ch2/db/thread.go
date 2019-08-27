package db

import (
	"time"
	"log"
	"bytes"
	"fmt"
)

/**
保存所有帖子相关代码
 */

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	Userid    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

const TimeFormat  = "2006-01-02 15:04:05"

func DateTimeNow() string{
	return time.Now().Format(TimeFormat)
}

/**
将time转为字符串输出
 */
func DateTimeToString(t time.Time) string{
	return t.Format(TimeFormat)
}

func (t *Thread) NumReplies()(count int){
	rows,err := Db.Query("select count(*) from posts where thread_id = ? ",t.Id)
	if err != nil{
		return
	}

	for rows.Next(){
		if err = rows.Scan(&count);err != nil{
			return
		}
	}

	rows.Close()
	return
}

func (t *Thread) Posts()(posts []Post,err error){
	s := "select id,uuid,body,user_id,thread_id,created_at from posts where thread_id = ?"
	rows,err := Db.Query(s,t.Id)
	if err != nil{
		log.Printf("query error[%s]/t:%v\n",s,err)
		return
	}

	for rows.Next(){
		post := Post{}
		if err = rows.Scan(&post.Id,&post.Uuid,&post.Body,&post.UserId,&post.ThreadId,&post.CreatedAt);err != nil{
			log.Printf("read data error:%v\n",err)
			return
		}
		posts = append(posts,post)
	}
	rows.Close()
	return
}

func (t *Thread) User()(user User){
	user = User{}
	err := Db.QueryRow("select id,uuid,name,email,created_at from users where id = ?",t.Userid).
		Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)
	if err != nil{
		return
	}

	return
}

func (t *Thread) CreatedAtDate() string{
	return DateTimeToString(t.CreatedAt)
}

func (t Thread) String() string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Thread[Id:%d,Uuid:%s,Topic:%s,Userid:%d,CreateAt:%v]\n", t.Id, t.Uuid, t.Topic, t.Userid, t.CreatedAt)
	return buf.String()
}

/**
从数据库里面取出所有帖子并将其返回给调用方
 */
func Threads() (threads []Thread, err error) {
	//s := "SELECT a.id, a.uuid, a.topic, b.`name`, a.created_at FROM threads a , users b where a.user_id = b.id ORDER BY a.created_at DESC"
	s := "SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC"

	if Db == nil {
		fmt.Printf("nil database connection\n")
		return
	}

	rows, err := Db.Query(s)
	if err != nil {
		log.Printf("query failed with sql[%s]:%v\n", s, err)
		return
	}


	for rows.Next() {
		//构造一个新对象
		th := Thread{}

		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.Userid, &th.CreatedAt); err != nil {
			log.Printf("read data error:%v\n", err)
			return
		}

		threads = append(threads, th)
	}
	rows.Close()
	return threads, nil
}


//get thread by uuid
func ThreadByUuid(uuid string)(conv Thread,err error){
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?",uuid).
		Scan(&conv.Id,&conv.Uuid,&conv.Topic,&conv.Userid,&conv.CreatedAt)
	return
}

func (post *Post) String()string{
	var buf bytes.Buffer
	fmt.Fprintf(&buf,"Post[id:%d\t,Uuid:%s\t,Body:%s\t,UserId:%d\t,ThreadId:%d\t,CreatedAt:%v]",
		post.Id,post.Uuid,post.Body,post.UserId,post.ThreadId,post.CreatedAt)
	return buf.String()
}

func (post *Post) User()(user User){
	user = User{}
	Db.QueryRow("select id,uuid,name,email,created_at from users where id = ?",post.UserId).
		Scan(&user.Id,&user.Uuid,&user.Name,&user.Email,&user.CreatedAt)
	log.Printf("%s\n",&user)
	return
}

func (post *Post) CreatedAtDate() string{
	return DateTimeToString(post.CreatedAt)
}