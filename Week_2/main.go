package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func errWarp(uid int) (string, error) {
	//数据库连接
	db, err := sql.Open("mysql", "root:123456Aa@tcp(127.0.0.1:3306)/geekdb?charset=utf8")
	db.Ping()
	if err != nil {
		fmt.Errorf("访问数据库失败，请检查！")
	}
	defer db.Close()

	var (
		//id int
		name string
	)

	err = db.QueryRow("select name from tb_user where id = ?", uid).Scan(&name)

	if err != nil {
		fmt.Println("query error")
		if err == sql.ErrNoRows {
			//log.Println("sql.ErrNoRows")
			return "", fmt.Errorf("%w, 数据为空", err)
		}
		log.Fatal(err)
	}
	return "", err
}

func main() {
	_, err := errWarp(100)
	err = errors.Unwrap(err)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("%s", err)
	}

}
