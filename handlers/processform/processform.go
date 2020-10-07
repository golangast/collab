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
tables.  This way we can create the dashboard
so that can show that user's
pages and keep the ID as a reference to
update the page table in that view
*/

func Processingform(c echo.Context) error {
	//get form data
	email := c.FormValue("email")
	pass := c.FormValue("password")
	var err error

	//populate that data in User struct
	user := users.User{Email: email, Password: pass}

	//regrab all users
	DBallusers := allusers.GetAllUsers()

	/*the for loop takes all users and looks for the
	email and compares it so that we can have a unique
	identifier on that user row and use it to send to the
	dashboard.  This way that user opens "their" dashboard.
	if the user is not found then it creates the user.*/
	for _, thisuser := range DBallusers {

		//check if emails match
		if thisuser.Email == email {
			//print stuff
			fmt.Println("email match for user ", thisuser)
			spew.Dump(thisuser.ID)

			//send that user to the template to get context
			return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
				"User": thisuser,
			}) //Render

		} else {
			fmt.Println("email didnt match, so creating user")
			//add data to database
			check := users.User.CreateUser(user, user.Email, user.Password)
			if check != true {
				fmt.Println("user not created")
			}
			return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
				"User": user,
			}) //Render
		} //if
	} //for

	return err
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
