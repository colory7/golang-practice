package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"os"
)

func main() {
	db, err := sql.Open("godror", `user="oracle" password="oracle" connectString="127.0.0.1:1521/XE"`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("oracle连接失败")
		os.Exit(2)
	}
	fmt.Println("oracle连接成功")
}
