package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init() "mysql"在database/sql中注册
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql"
	db, err := sql.Open("mysql", dsn)
	//sql.Open()不会校验用户名和密码是否正确，只会检测dsn的格式是否正确
	if err != nil {
		fmt.Printf("dsn %s invalid, err:%v\n", dsn, err)
		return
	}

	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s fialed, err:%v", dsn, err)
		return
	}
}
