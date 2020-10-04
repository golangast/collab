package form

import (
	"net/http"

	"github.com/labstack/echo"
)

func Form(c echo.Context) error {

	return c.Render(http.StatusOK, "form.html", map[string]interface{}{
		"User": "none",
	})

}
