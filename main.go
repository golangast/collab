package main

import (
	g "collab/db/getallusers"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	u := g.GetAllUsers()
	fmt.Println(u)
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "The beggining...")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
