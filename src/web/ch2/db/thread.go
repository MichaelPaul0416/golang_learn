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

func (t Thread) String() string {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Thread[Id:%d,Uuid:%s,Topic:%s,Userid:%d,CreateAt:%v]\n", t.Id, t.Uuid, t.Topic, t.Userid, t.CreatedAt)
	return buf.String()
}


/**
从数据库里面取出所有帖子并将其返回给调用方
 */
func Threads() (threads []Thread, err error) {
	s := "select id,uuid,topic,user_id,created_at from threads order by created_at desc"

	if Db == nil{
		fmt.Printf("nil database connection\n")
		return
	}

	rows, err := Db.Query(s)
	if err != nil {
		log.Printf("query failed with sql[%s]:%v\n", s, err)
		return
	}

	//var ts []Thread
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
