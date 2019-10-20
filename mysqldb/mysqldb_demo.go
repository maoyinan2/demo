package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/* tips
1, NullBool, Scan | Value | Valid
2, DB | Result（结果状态） | Row(s)（结果集） | Stmt | Tx
3,
*/
func main() {

	/*	dataSource: [user:password]@[protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		paramN : https://github.com/go-sql-driver/mysql
		eg: user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true	*/
	// ? config.json
	db, err := sql.Open("mysql", "root:888@/test")
	defer db.Close()
	checkErr(err)

	if err = db.Ping(); err != nil { // db是否实际可用，密码错误sql.Open不会报错
		fmt.Println("db ping error!")
	}

	/* DB: Exec | Query | QueryRow | Prepare | Begin
	Prepare server返回Stmt（绑定连接池中空闲连接）
	*/
	stmt, err := db.Prepare("insert userinfo set username = ?,department=?,created=?")
	checkErr(err)

	//执行准备好的Stmt
	res, err := stmt.Exec("user1", "computing", "2019-02-20")
	checkErr(err)

	//获取上一个，即上面insert操作的ID
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id) //1

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("user1update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect) //1

	//查询数据
	var s sql.NullString
	err = db.QueryRow("select username from userinfo where uid=?", 1).Scan(&s)
	if s.Valid {
		fmt.Println("s.valid.")
	} else {
		fmt.Println("s is null.")
	}

	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	var uid int
	var username, department string
	var created []uint8

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created) //1 user1update computing 2019-02-20
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}
	defer rows.Close()

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect) //1

}
