package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func findByPk(pk int) int {
	var num int = 0
	db, err := sql.Open("mysql", "root:@123456(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtOut, err := db.Prepare("show tables")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	err = stmtOut.QueryRow(pk).Scan(&num)
	if err != nil {
		panic(err.Error())
	}
	return num
}
