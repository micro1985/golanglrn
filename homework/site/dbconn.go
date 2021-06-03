package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	//	fmt.Println("Work with MySQL")

	db, err := sql.Open("mysql", "gouser:123456@tcp(192.168.40.22:3306)/godb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/*	insert, err := db.Query("INSERT INTO `gotable` (`name`, `age`) VALUES('Bob2', 33)")
		if err != nil {
			panic(err)
		}
		defer insert.Close()
	*/
	res, err := db.Query("SELECT * FROM `gotable`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with Age %d", user.Name, user.Age))
	}

	defer res.Close()

	fmt.Println("Connect to MySQL done")
}
