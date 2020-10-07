package pageedit

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Pageedit(c echo.Context) error {
	id := c.Param("id")
	pageid := c.Param("pageid")

	type URL struct {
		Id     string
		Pageid string
	}

	u := URL{Id: id, Pageid: pageid}
	fmt.Println(u.Id)
	return c.String(http.StatusOK, u.Id)
}
