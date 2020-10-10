package main

import (
	//u "collab/pkg/user"
	c "collab/handlers/contentedit"
	f "collab/handlers/form"
	pc "collab/handlers/pagecreate"
	p "collab/handlers/processform"
	"io"
	"net/http"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/labstack/echo"
)

/*database layout
database name is db
   user     | post | page
  -id       -id       -id
  -email              -url
  -password -content  -title
			-date     -content
*/
func main() {
	// Server header
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e := echo.New()

	e.Renderer = renderer

	e.GET("/form", f.Form)
	e.POST("/dashboard", p.Processingform)
	e.GET("/page-create/:userid", pc.Pagecreate)
	e.GET("/contentedit", c.Contentedit)
	e.Static("/static", "assets")
	e.Static("/", "assets")
	// Debug/middleware/templates
	e.Debug = true
	s := NewStats()
	e.Use(s.Process)

	e.Logger.Fatal(e.Start(":1323"))

}

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

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
