package main

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
)

// User .
type User struct {
	UID  int64
	Name string
}

func main() {
	/* connection string
	   url-> sqlserver://sa:mypass@localhost:1234?database=master&connection+timeout=30
	   ado-> server=localhost;port=1433;user id=sa;password=123;database=master;connection timeout=5(s);encrypt=false
	   server=localhost;user id=sa;database=test;app name=go-mssqldb
	*/
	var (
		server   = "localhost"
		user     = "sa"
		password = "myn123456"
		database = "test"
		port     = 1433
	)

	/* error
	Login error: read tcp 10.1.8.159:5902->192.168.7.209:1433: wsarecv: An existing connection was forcibly closed by the remote host
	try => encrypt=disable */
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable", server, user, password, port, database)
	engine, err := xorm.NewEngine("mssql", connString)
	// conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println("Open connection failed:", err.Error())
	}
	// defer conn.Close()
	var affected int64
	user1 := User{Name: "maoyinan"}
	affected, err = engine.Insert(&user1)
	fmt.Println(affected)
	// err := engine.Sync2(new(User))

	// results, err := engine.Query("select * from userinfo")

	// if err = conn.Ping(); err != nil { // db是否实际可用，密码错误sql.Open不会报错
	// 	fmt.Println("conn ping error!", err.Error())
	// }

	// stmt, err := conn.Prepare("select 1, 'abc'")
	// if err != nil {
	// 	log.Fatal("Prepare failed:", err.Error())
	// }
	// defer stmt.Close()

	// row := stmt.QueryRow()
	// var somenumber int64
	// var somechars string
	// err = row.Scan(&somenumber, &somechars)
	// if err != nil {
	// 	log.Fatal("Scan failed:", err.Error())
	// }
	// fmt.Printf("somenumber:%d\n", somenumber)
	// fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")
	/* data type */

	// fmt.Println("hello, Go!")
}
