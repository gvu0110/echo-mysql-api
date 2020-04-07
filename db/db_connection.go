package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/foo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("db is connected")
	return db
}
