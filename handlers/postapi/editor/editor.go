package editor

import (
	p "collab/pkg/post"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func Editor(c echo.Context) error {
	//URL DATA
	pageid := c.Param("pageid")
	type URL struct {
		Pageid string
	}
	pageids := URL{Pageid: pageid}

	//FORM DATA
	html := c.FormValue("html")
	css := c.FormValue("css")
	js := c.FormValue("js")

	//initialize
	postcode := p.Post{Pageid: pageids.Pageid, HTML: html, CSS: css, JS: js}

	//SAVE IN DATABASE
	check := p.Post.CreatePost(postcode, postcode.Pageid, postcode.HTML, postcode.CSS, postcode.JS)
	if check != true {
		fmt.Println("post not created")
	}

	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"User": user,
	}) //Render

}
