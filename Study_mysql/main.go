package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

const (
	USERNAME = "xieli"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "test"
)

var Db *sqlx.DB

// https://pkg.go.dev/github.com/jmoiron/sqlx@v1.3.4#readme-usage
func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	// Db, err := sqlx.Open("mysql", dsn) // 注意这里有问题，如果是":="，并没有赋值到外部声明的Db
	database, err := sqlx.Open("mysql", dsn)
	// database, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("Open mysql failed, err: ", err)
		return
	}
	// 传到外部
	Db = database
	fmt.Println("Open mysql ok")
}

func main() {
	defer Db.Close()
	// insertOperate()
	queryOperate()
}

// 插入操作
func insertOperate() {
	result, err := Db.Exec("insert INTO person(username, sex, email) values(?,?,?)", "zhangsan", "man", "zhangsan@qq.com")
	if err != nil {
		fmt.Println("Exec failed, ", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed, ", err)
		return
	}
	fmt.Println("Insert ok: ", id)
}

// 查询操作
func queryOperate() {
	// var person []Person
	person := make([]Person, 10)
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 3)
	if err != nil {
		fmt.Println("Select failed, err: ", err)
		return
	}
	fmt.Println("Select ok: ", person)
}
