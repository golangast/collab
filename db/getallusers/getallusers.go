package getallusers

import (
	dbconn "collab/db"
	u "collab/pkg/user"
	"fmt"
)

//GetAllUsers get all users
func GetAllUsers() []u.User {
	//opening database
	data := dbconn.Conn()
	//data from database to put into User struct

	var (
		id    int
		email string
		pass  string
		user  []u.User
	)
	i := 0
	//get from database
	rows, err := data.Query("select * from user")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &email, &pass)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}

		user = append(user, u.User{ID: id, Email: email, Password: pass})

	}
	defer rows.Close()
	defer data.Close()
	//fmt.Println(user)
	return user
}
