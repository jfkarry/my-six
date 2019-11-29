package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(127.0.0.1)/t1?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败")
	} else {
		fmt.Println("连接数据库成功")
	}
	insertDB(db)
	deleteDB(db)
	updateDB(db)
	selectDB(db)
	defer db.Close()
}
func insertDB(db *sql.DB) {
	stmt, err := db.Prepare("insert into t11(name) values (?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec("cca")

}
func deleteDB(db *sql.DB) {
	stmt, err := db.Prepare("delete from t11 where id = 2")
	if err != nil {
		log.Fatal(err);
	}
	stmt.Exec();
}
func updateDB(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE t11 SET name = 'ttt' WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec();
}
func selectDB(db *sql.DB) {
	stmt, err := db.Query("select * from t11;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next() {
		var id int
		var name string
		//var name sql.NullString
		err := stmt.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}
