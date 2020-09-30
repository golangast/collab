package post

import (
	dbconn "collab/db"
	"fmt"
	"log"
)

//User for database
type Post struct {
	ID      int
	Email   string
	Content string
	Date    string
}

//CreateUser creates a user
func (p Post) CreatePost(e string, c string, d string) bool {
	//opening database
	data := dbconn.Conn()
	// query
	stmt, err := data.Prepare("INSERT INTO post(email, content, date) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//initializing User
	userstemp := Post{Email: e, Content: c, Date: d}
	//checking whats going into database because res below prints way too much info
	fmt.Println(userstemp, "- added to database")
	//add to the database
	res, err := stmt.Exec(userstemp.Email, userstemp.Content, userstemp.Date)
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
