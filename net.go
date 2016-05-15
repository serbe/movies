package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type tmpl struct {
    templates *template.Template
}

func initRender() *tmpl {

	t := &tmpl{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	// r := multitemplate.New()

	// r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	// r.AddFromFiles("react", "templates/react.html")

	return t
}

// Render html template
func (t *tmpl) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func (app *application) initServer() {
	e := echo.New()

	e.SetRenderer(initRender())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", root)

	e.GET("/movie", app.getOneMovieJSON)
	e.GET("/movies", app.getMoviesJSON)

	e.Static("/public", "public")
	e.File("/favicon.ico", "public/favicon.ico")

	e.Run(standard.New(":" + app.config.Web.Port))

	app.server = e
}
