package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/demo?charsetutf8`)
	if err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}
	if err := DB.Ping(); err != nil {
		fmt.Println("数据库请求失败")
		os.Exit(1)
	}
	fmt.Println("数据库连接成功")
}
