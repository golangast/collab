package contentedit

import (
	"net/http"

	"github.com/labstack/echo"
)

func Contentedit(c echo.Context) error {
	userId := c.Param("userid")

	type URL struct {
		UserId string
	}

	u := URL{UserId: userId}

	return c.Render(http.StatusOK, "contentedit.html", map[string]interface{}{
		"User": u,
	}) //Render

}
