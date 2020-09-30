package user

import (
	dbconn "collab/db"
	"fmt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) CreateUser(e string, p string) bool {
	//opening database
	data := dbconn.Conn()
	// query
	stmt, err := data.Prepare("INSERT INTO user(email, password) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	userstemp := User{Email: e, Password: p}
	fmt.Println(userstemp)

	res, err := stmt.Exec(userstemp.Email, userstemp.Password)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return true
}
