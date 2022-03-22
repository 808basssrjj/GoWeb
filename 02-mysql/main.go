package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

// initMysql 初始化连接
func initMysql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to mysql err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(200) // 最大连接数
	db.SetMaxIdleConns(10)  // 最大空闲连接数
	return nil
}

func main() {
	if err := initMysql(); err != nil {
		fmt.Printf("connect to mysql err:%v\n", err)
		return
	}
	defer db.Close()
	fmt.Println("connect sucess")

	queryRow()
	queryMul()
	insert()
	update()
	deleteRow()
}

type user struct {
	id   int
	age  int
	name string
}

// 预处理
// 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
// 避免SQL注入问题
// prepareQuery 预处理查询示例
func prepareQuery() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}


// queryRow 查询一条
func queryRow() {
	sqlStr := "select id, name, age from user where id = ?"
	var u user

	row := db.QueryRow(sqlStr, 1)
	// 必须使用scan方法，不然会一直占用连接
	err := row.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

func queryMul() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed err:%v\n", err)
	}
	// 关闭rows 释放持有的数据库连接
	defer rows.Close()

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.id, &u.name, &u.age); err != nil {
			fmt.Printf("scan failed err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// insert 插入数据
func insert() {
	sqlStr := "insert into user (name, age) values (?, ?)"
	res, err := db.Exec(sqlStr, "娜美", 18)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	theID, err := res.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// update 更新数据
func update() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// deleteRow 删除数据
func deleteRow() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
