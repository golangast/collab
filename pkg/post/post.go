package post

import (
	dbconn "collab/db"
	"fmt"
	"log"
)

type Post struct {
	Pageid string
	HTML   string `json:"html" form:"html" query:"html"`
	CSS    string `json:"css" form:"css" query:"css"`
	JS     string `json:"js" form:"js" query:"js"`
}

//CreateUser creates a user
func (p Post) CreatePost(pi string, h string, c string, j string) bool {
	//opening database
	data := dbconn.Conn()
	// query
	stmt, err := data.Prepare("INSERT INTO post(pageid, html, css, js) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//initializing Post
	posttemp := Post{Pageid: pi, HTML: h, CSS: c, JS: j}
	//checking whats going into database because res below prints way too much info
	fmt.Println(posttemp, "- added to database")
	//add to the database
	res, err := stmt.Exec(posttemp.Pageid, posttemp.HTML, posttemp.CSS, posttemp.JS)
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
