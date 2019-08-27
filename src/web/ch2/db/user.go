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

func (user *User) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "User[Id:%d\t,Uuid:%s\t,Name:%s\t,Email:%s\t,Password:%s\t,CreatedAt:%v]", user.Id, user.Uuid, user.Name, user.Email, user.Password, user.CreatedAt)
	return buf.String()
}

//create a session for current user
func (user *User) CreateSession() (s Session, err error) {
	statement := "insert into sessions (uuid,email,user_id,created_at) values (?,?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Printf("create error[%s]\t:%v\n", statement, err)
		return
	}

	defer stmt.Close()

	session := Session{}

	uuid := CreateUUID()
	session.Uuid = uuid
	session.UserId = user.Id
	session.Email = user.Email
	session.CreatedAt = time.Now()
	res, err := stmt.Exec(uuid, user.Email, user.Id, session.CreatedAt)

	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	session.Id = int(id)
	return session, nil
}

// get a session from user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where user_id=?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// check is session is valid
//方法绑定的对象的形参s 在一个包中最好保持一致
func (s *Session) CheckValid() (valid bool, err error) {
	err = Db.QueryRow("select id,uuid,email,user_id,created_at from sessions where uuid = ?", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)

	if err != nil {
		valid = false
		return
	}

	if s.Id != 0 {
		valid = true
	}
	return
}

//delete by uuid for session
func (s *Session) DeleteByUuid() (err error) {
	sql := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.Uuid)
	if err != nil {
		log.Printf("delete session by uuid error:%v\n", err)
		return err
	}

	return nil
}

//get the user from session
func (s *Session) User() (user User, err error) {
	//方法内部也可以直接使用返回的出参
	user = User{}
	sql := "select id,uuid,name,email,created_at from users where id = ?"
	//查询一行
	err = Db.QueryRow(sql,s.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)

	return
}

func SessionDeleteAll() (err error) {
	//这里使用=赋值，而不是:=赋值，因为err使用的是返回的形参err
	_, err = Db.Exec("delete from sessions")
	return
}

// create user pojo
func (user *User) Create() (id int64, err error) {
	sql := "insert into users (uuid,name,email,password,created_at) value (?,?,?,?,?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}

	res, err := stmt.Exec(CreateUUID(), user.Name, user.Email, Encrypt(user.Password), DateTimeNow())
	if err != nil {
		return
	}
	uid, err := res.LastInsertId()
	user.Id = int(uid)
	return uid, err
}

//delete user
func (user *User) Delete() (err error) {
	sql := "delete from users where id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	return
}

//update user info
func (user *User) Update() (err error) {
	sql := "update user s set name=?,email=? where id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = Db.Exec(user.Name, user.Email, user.Id)
	if err != nil {
		return err
	}
	return
}

func (user *User) CreateThread(topic string) (conv Thread, err error) {
	sql := "insert into threads (uuid,topic,user_id,created_at) values(?,?,?,?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}

	defer stmt.Close()
	(&conv).Uuid = CreateUUID()
	(&conv).CreatedAt = time.Now()
	(&conv).Topic = topic
	(&conv).Userid = user.Id

	res, err := stmt.Exec(conv.Uuid, conv.Topic, conv.Userid, conv.CreatedAt)
	if err != nil {
		return
	}

	tid, err := res.LastInsertId()
	(&conv).Id = int(tid)
	return
}

func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	sql := "insert into posts (uuid,body,user_id,thread_id,created_at) values (?,?,?,?,?)"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	var p Post
	(&p).CreatedAt = time.Now()
	(&p).Uuid = CreateUUID()
	(&p).Body = body
	(&p).ThreadId = thread.Id
	(&p).UserId = user.Id

	res, err := stmt.Exec(p.Uuid, p.Body, p.UserId, p.ThreadId, p.CreatedAt)
	if err != nil {
		return
	}
	pid, err := res.LastInsertId()
	if err != nil {
		return
	}
	(&p).Id = int(pid)
	return p, nil
}

func UserDeleteAll() (err error) {
	_, err = Db.Exec("delete from users")
	return
}

func Users() (users []User, err error) {
	rows, err := Db.Query("select id,uuid,name,email,password,created_at from users");
	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}

		users = append(users, user)
	}

	rows.Close()
	return
}

func UserByEmail(email string) (user User, err error) {
	sql := "select id,uuid,name,email,password,created_at from users where email = ?"
	if err != nil{
		return
	}
	err = Db.QueryRow(sql,email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByUuid(uuid string) (user User, err error) {
	sql := "select id,uuid,name,email,password,created_at from users where uuid = ?"
	err = Db.QueryRow(sql,uuid).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
