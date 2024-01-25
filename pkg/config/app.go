package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect(){
	d, err := sql.Open("mysql","root:my-secret-pw@/tab?parseTime=true")
	if err != nil {
		panic(err)
	}

	fmt.Println("pinging db...")
	err = d.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db")

	db = d
}

func GetDB() *sql.DB {
	return db
}