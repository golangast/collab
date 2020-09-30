package main

import (
	u "collab/pkg/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	//checking the CreateUser works
	user := u.User{Email: "zach@yahoo.com", Password: "qw"}
	check := u.User.CreateUser(user, user.Email, user.Password)
	fmt.Println(check)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "The beggining...")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
