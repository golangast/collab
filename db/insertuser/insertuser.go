/*Currently this would not need the sanitizer
but as a template it is good to have for any input*/

package insertuser

import (
	dbconn "collab/db"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

/*The Data*/
//User not imported because of sanitizer interface
type User struct {
	ID       int
	Email    string
	Password string
}

//Data is a wrapper of User/Sanitizer
type Data struct {
	U User
	S Sanitizer
}

/*Insertuser is the actual function thats going to be exported*/
func Insertuser(w http.ResponseWriter, r *http.Request) {

	/*not needed currently but if we make a login
	for the user then this can be used*/
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	unmarshbody, err := UnmarshalLogin(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	//opening database
	data := dbconn.Conn()

	// query
	stmt, err := data.Prepare("INSERT INTO user(email, pass) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	userstemp := Data{U: User{Email: unmarshbody.Email, Password: unmarshbody.Password}}
	fmt.Println(userstemp)

	u := userstemp
	s, err := Save(u)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(s.U.Email, s.U.Password)
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
	fmt.Println("reached query")

	w.WriteHeader(http.StatusOK)
}

//This is where the sanitizing begins.
//Can be used for a login or any input
func UnmarshalLogin(data []byte) (User, error) {
	var u User
	fmt.Print("starting unmarshal")
	err := json.Unmarshal(data, &u)
	fmt.Print("is starting", data)

	return u, err
}

//Runs the Go's escapeString
func (p User) Sanitize() {
	p.Email = html.EscapeString(p.Email)
	p.Password = html.EscapeString(p.Password)
}

type Sanitizer interface {
	Sanitize()
}

//If the sanitize makes it then no error
func Save(u Data) (Data, error) {
	var err error

	// type assertion for Sanitizer (could also use a type switch)
	s, ok := u.S.(Sanitizer)

	if !ok {
		if err != nil {
			log.Fatal(err)
		}
		// ... save without sanitization
		return u, err
	}

	s.Sanitize()
	return u, err
}
