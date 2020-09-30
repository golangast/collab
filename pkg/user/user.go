package user

import (
	dbconn "collab/db"
	"log"
)

//User for database
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//CreateUser creates a user
func (u User) CreateUser(e string, p string) bool {
	//opening database
	data := dbconn.Conn()
	// query
	stmt, err := data.Prepare("INSERT INTO user(email, password) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	//initializing User
	userstemp := User{Email: e, Password: p}

	//add to the database
	res, err := stmt.Exec(userstemp.Email, userstemp.Password)
	if err != nil {
		log.Fatal(err)
	}
	//if error then print first and last id
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
