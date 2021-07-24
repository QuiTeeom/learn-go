package main

import (
	"database/sql"
	errors2 "errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func main() {
	db := openDB()
	defer db.Close()

	initDB(db)

	if user, err := getUser(db); err != nil {
		if errors2.Is(err, sql.ErrNoRows) {
			fmt.Printf("查询不到用户 %+v", err)
		} else {
			fmt.Println(err)
			panic(err)
		}
	} else {
		fmt.Println(user)
	}

}

func getUser(db *sql.DB) (*User, error) {
	scans := make([]interface{}, 2)
	for i := range scans {
		scans[i] = &scans[i]
	}
	query := "SELECT * FROM user"

	err := db.QueryRow(query).Scan(scans...)
	if err != nil {
		return nil, errors.Wrapf(err, "sql:%s", query)
	} else {
		return &User{
			Id:   scans[0].(string),
			Name: scans[1].(string),
		}, nil
	}
}

func openDB() *sql.DB {
	if db, err := sql.Open("sqlite3", "maojian-homeworks/week2/demo.db"); err != nil {
		// demo程序基础，所以打开db出错就直接退出
		panic(err)
	} else {
		return db
	}
}

func initDB(db *sql.DB) {
	stmt := `
	CREATE TABLE IF NOT EXISTS user(id varchar ,name varchar ) 
	`
	_, err := db.Exec(stmt)
	if err != nil {
		// demo程序基础，所以打开db出错就直接退出
		panic(err)
	}

}

type User struct {
	Id   string
	Name string
}
