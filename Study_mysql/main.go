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
	createTable()
	// insertOperate()
	// queryOperate()
	// updateOperate()
	// deleteOperate()
	// transactionOperate()
}

func createTable() {
	// 创建表名为place的表
	sql := `
		create table if not exists place(
			id int(11) primary key auto_increment,
			country varchar(200),
			city varchar(200),
			telcode int
		)ENGINE=InnoDB default charset=utf8;
	`
	if _, err := Db.Exec(sql); err != nil {
		fmt.Println("Create table failed, err: ", err)
	}
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
	fmt.Println("Insert ok, affected rows: ", id)
}

// 查询操作
func queryOperate() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("Select failed, err: ", err)
		return
	}
	fmt.Println("Select ok: ", person)
}

// 更改操作
func updateOperate() {
	// 在一条UPDATE语句中，如果要更新多个字段，字段间不能使用“AND”，而应该用逗号分隔
	result, err := Db.Exec("update person set username=?, email=? where user_id=?", "lisi", "lisi@qq.com", 4)
	if err != nil {
		fmt.Println("Update failed, err: ", err)
		return
	}
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed, err:", err)
		return
	}
	fmt.Println("Update ok, affected rows: ", row)
}

// 删除操作
func deleteOperate() {
	result, err := Db.Exec("delete from person where user_id=?", 3)
	if err != nil {
		fmt.Println("Delete failed, err: ", err)
		return
	}
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed, err:", err)
		return
	}
	fmt.Println("Delete ok, affected rows: ", row)

}

// 事务操作
func transactionOperate() {
	// 开始事务
	tx, err := Db.Begin()
	if err != nil {
		fmt.Println("Begin failed, err: ", err)
		return
	}
	result, err := tx.Exec("insert into person(username, sex, email) values(?,?,?)", "wangwu", "man", "wangwu@qq.com")
	if err != nil {
		fmt.Println("Exec failed, err:", err)
		// 事务回滚
		tx.Rollback()
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed, err:", err)
		// 事务回滚
		tx.Rollback()
		return
	}
	fmt.Println("Insert ok, id: ", id)

	result, err = tx.Exec("update person set sex=? where username=?", "women", "lisi")
	if err != nil {
		fmt.Println("Update failed, err:", err)
		tx.Rollback()
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed, err:", err)
		// 事务回滚
		tx.Rollback()
		return
	}
	fmt.Println("Update ok, affected rows:", rows)

	// 事务确认
	tx.Commit()
}
