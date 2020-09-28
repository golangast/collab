package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//start connection to db
func Conn() *sql.DB {
	fmt.Println("db begin")
	db, err := sql.Open("mysql", "root:@/db")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("open ")
	}

	//check if it pings
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping ")
	}
	return db
}
