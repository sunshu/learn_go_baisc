package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	//mysql data source name
	Dsn string
}

func main() {
	dbw := DbWorker{
		Dsn: "root:root@tcp(123.206.95.89:3306)/user",
	}
	db, err := sql.Open("mysql",

		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("打开成功")

	defer db.Close()
}
