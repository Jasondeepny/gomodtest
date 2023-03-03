package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func DbTest() {
	reps, err := Db.Exec("insert into person(username,sex,email)values(?,?,?)", "stu001", "man", "stu001@edu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Db.Close()

	id, err := reps.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("insert success,id :", id)
}

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(10.1.11.134:3306)/test")
	if err != nil {
		fmt.Println("open mysql error", err)
		return
	}
	Db = database
}

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}
type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}
