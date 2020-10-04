package processform

import (
	allusers "collab/db/getallusers"
	users "collab/pkg/user"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

/*
The point of the method is so that you can
process the request from the form and then
get the context of that user in the datbase
tables.  This way we can create the URI/URL
so that the dashboard can show that user's
pages and keep the ID as a reference to
update the page table in that view
*/

func Processingform(c echo.Context) error {
	//get form data
	email := c.FormValue("email")
	pass := c.FormValue("password")
	//populate that data in User strcut
	user := users.User{Email: email, Password: pass}
	//add data to database
	check := users.User.CreateUser(user, user.Email, user.Password)
	if check != true {
		fmt.Println("user not created")
	}
	//regrab all users
	ausers := allusers.GetAllUsers()
	//compare them to get the email
	for _, value := range ausers {
		//check if emails match
		if value.Email == email {
			fmt.Println("email match!")
			spew.Dump(value.ID)
			//send that user to the template to get context
			return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
				"User": value,
			}) //Render
		} else {
			fmt.Println("email didnt match")
		} //if
	} //for
	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"User": user,
	}) //Render
} //Processingform

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
