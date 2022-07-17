package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ydg941020/geetTime/work1/dao"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/user?charset=utf8")

	CreateTable := "CREATE TABLE user(" +
		"id INT(10) NOT NULL AUTO_INCREMENT," +
		"name VARCHAR(64) NULL DEFAULT NULL," +
		"age INT(10) DEFAULT NULL,PRIMARY KEY (id))" +
		"ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	db.Exec(CreateTable)
}

func main() {
	user, err := dao.QueryMultiRowDemo(db)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	fmt.Printf("user:%v", user)
}
