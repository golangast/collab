package userpages

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetPages(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)

}
