package db

import (
	"time"
	"bytes"
	"fmt"
	"log"
)

/**
用户操作相关
 */

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (s *Session) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Session[Id:%d\t,Uuid:%s\t,Email:%s\t,UserId:%d\t,CreatedAt:%v]", s.Id, s.Uuid, s.Email, s.UserId, s.CreatedAt)
	return buf.String()
}

//create a session for current user
func (user *User) CreateSession() (s Session, err error) {
	statement := "insert into sessions (uuid,email,user_id,created_at) values (?,?,?,?)"
	stmt,err := Db.Prepare(statement)
	if err != nil{
		log.Printf("create error[%s]\t:%v\n",statement,err)
		return
	}

	defer stmt.Close()

	session := Session{}

	uuid := CreateUUID()
	session.Uuid = uuid
	session.UserId = user.Id
	session.Email = user.Email
	session.CreatedAt = time.Now()
	res,err := stmt.Exec(uuid,user.Email,user.Id,session.CreatedAt)

	if err != nil{
		return
	}

	id,err := res.LastInsertId()
	session.Id = int(id)
	return session,nil
}

// get a session from user
func (user *User) Session(session Session,err error){
	session = Session{}
	err = Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where user_id=?",user.Id).
		Scan(&session.Id,&session.Uuid,&session.Email,&session.UserId,&session.CreatedAt)
	return
}

// check is session is valid
func (session *Session) CheckValid(valid bool,err error){
	err = Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where uuid = ?",session.Uuid).
		Scan(&session.Id,&session.Uuid,&session.Email,&session.Uuid,&session.CreatedAt)

	if err != nil{
		valid = false
		return
	}

	if session.Id != 0{
		valid = true
	}
	return
}