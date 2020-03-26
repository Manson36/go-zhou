package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB //为了在其他函数中可以使用db，需要做一个全局的变量
func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql"
	db, err = sql.Open("mysql", dsn) //这里的db是一个链接池
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println("init DB error :", err)
	}
	fmt.Println("连接数据库成功")
	SelectById()
}

type user struct {
	id   int
	name string
	age  int
}

func SelectById() {
	var u user
	//1.写查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id = ?`
	//2.执行
	rowObj := db.QueryRow(sqlStr, 1) //从连接池中拿出一个连接去数据库查询单条记录
	//3.拿到结果
	rowObj.Scan(&u.id, &u.name, &u.age) //因为要改变u的值，所以使用指针；Scan方法中有释放连接的操作
	//4.打印结果
	fmt.Printf("u1:%#v\n", u)
}

func SelectMore(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v", err)
		return
	}
	//非常重要，关闭rows释放连接
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println("Scan failed, err:", err.Error())
			return
		}
		fmt.Printf("id:%d name:%s age:%d", u.id, u.name, u.age)
	}
}

func Insert() {
	//1.sql语句
	sqlStr := `insert into user(name, age) values("李沁", 29)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("insert failed; err:", err.Error())
		return
	}
	//如果是插入操作，能拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("get id failed, err:", err.Error())
		return
	}
	fmt.Println("id:", id)
}

func Update() {
	sqlStr := `update user set age = ? where id = ?`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return
	}
	//更新操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get rowsAffected failed, err:", err)
		return
	}
	fmt.Println("update success, affected rows :", n)
}

func PrepareInsert() {
	sqlStr := `insert into user(name, age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed, err:", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"李沁": 29,
		"白鹿": 25,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func TransactionDemo() {
	//1.开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed, err:", err)
		return
	}
	//执行多个sql操作
	sqlStr1 := `update user set age = age -2 where id =1`
	sqlStr2 := `update user set age = age +2 where id =3`
	_, err = db.Exec(sqlStr1)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sql1错误，err:", err)
		return
	}
	_, err = db.Exec(sqlStr2)
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("执行sql2错误，err:", err)
		return
	}
	//sql语句都执行成功，提交本次事务
	err = tx.Commit()
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("提交出错了，err:", err)
		return
	}
	fmt.Println("事务提交成功")
}
