package pagecreate

import (
	"net/http"

	"github.com/labstack/echo"
)

func Pagecreate(c echo.Context) error {
	userId := c.Param("userid")

	type URL struct {
		UserId string
	}

	u := URL{UserId: userId}

	return c.Render(http.StatusOK, "pagecreate.html", map[string]interface{}{
		"User": u,
	}) //Render

}
