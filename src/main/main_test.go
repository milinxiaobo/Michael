package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlTest() {
	db, err := sql.Open("mysql", "root:linxiaobo@tcp(10.5.232.78:3306)/linxiaobo")
	if err != nil {
		fmt.Println(err)
		return
	}
	// db.Exec("show databases")
	// db.Exec("insert into test values (1, 1)")
	// tx, err := db.Begin()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// tx.Exec("select * from test")
	// tx.Commit()
	db.Exec("select * from test")
}
