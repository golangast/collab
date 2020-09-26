package Home

import (
	"fmt"
	"net/http"
	"text/template"

	Header "github.com/golangast/collab/go/Handlers/Header"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// tpl holds all parsed templates
var tpl *template.Template

// index serves the index html file
func Home(w http.ResponseWriter, r *http.Request) {
	Header.Headers(w, r)
	fmt.Println("index started ")

	if r.Method == "GET" {

		tpl.ExecuteTemplate(w, "index.html", nil)
	}

	if r.Method == "POST" {

		tpl.ExecuteTemplate(w, "index.html", nil)
	}

}
