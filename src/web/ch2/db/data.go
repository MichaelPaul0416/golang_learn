package db

import (
	"database/sql"
	"strings"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"crypto/sha1"
	"crypto/rand"
)

/**
与mysql数据相关连接
 */

var Db *sql.DB

const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "golang"
)

func init() {
	var err error
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	log.Printf("mysql driver:%s\n", path)
	//不能Db,err := sql.Open("mysql",path) 这样写
	//因为这样写的话，其实是在init方法体内重新定义了一个同名的局部变量Db
	Db, err = sql.Open("mysql", path)

	if err != nil {
		log.Fatal(err)
	}

	Db.SetMaxOpenConns(100)
	Db.SetMaxIdleConns(10)

	if err := Db.Ping(); err != nil {
		log.Printf("ping database failed\n")
		return
	}

	log.Printf("connect to database successfully:%t", Db == nil)
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
