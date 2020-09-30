package main

import (
	//u "collab/pkg/user"
	//p "collab/pkg/post"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//for testing createuser
	// user := u.User{Email: "zach@yahoo.com", Password: "qw"}
	// check := u.User.CreateUser(user, user.Email, user.Password)
	// fmt.Println(check)

	//for testing createpost
	// user := p.Post{Email: "zach@yahoo.com", Content: "this is content", Date: "33/33/3333"}
	// check := p.Post.CreatePost(user, user.Email, user.Content, user.Date)
	// fmt.Println(check)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "The beggining...")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
