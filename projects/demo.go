package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/backbay2")
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败！\n")
		return
	}
	//	1.执行数据库操作语句
	/*	sql := `insert into emp values(4,"444")`
		result, _ := db.Exec(sql)
		n, _ := result.RowsAffected()
		fmt.Println("受影响的记录数：", n)
		fmt.Printf("%T", result)*/
	//	2.执行预处理
	/*stu:=[2][2]string{{"5","555"},{"6","666"}}
	stmt,_:=db.Prepare("insert into emp values (?,?)")
	for i,s:=range stu {
		fmt.Println(i)
		fmt.Println(s)
		stmt.Exec(s[0],s[1])
	}*/
	//	3.单行查询
	/*	var id, name string
		rows := db.QueryRow("select * from emp where id=4")
		rows.Scan(&id, &name)
		fmt.Println(id, "--", name)*/
	//	4.多行查询
	var id, name string
	rows, err := db.Query("select * from emp")
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, "--", name)
	}

}
