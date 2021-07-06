package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func Connect() (db *sql.DB){
	dsn := "root:@tcp(localhost:3306)/crawler"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("ko connect")
		panic(err.Error())
	}else {
		fmt.Println("--------")
	}
	return db
}