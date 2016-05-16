package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"

	"github.com/dinever/golf"
)

type ctx struct {
	title string
}

func (app *application) initRender() {
	var templates map[string]*template.Template
	// var funcMap = template.FuncMap{"add": add, "nolast": nolast}
	index := []string{"templates/base.html", "templates/index.html"}
	templates["index"] = template.Must(template.New("base.html").ParseFiles(index...))
	// templates["index"] = template.Must(template.New("base.html").Funcs(funcMap).ParseFiles(index...))

	app.templates = templates
}

func (app *application) render(w http.ResponseWriter, c ctx, name string) error {
	tmpl, ok := app.templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}
	err := tmpl.ExecuteTemplate(w, "base.html", c)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	return err
}

func (app *application) initServer() {
	g := golf.New()


	mux := http.NewServeMux()
  	mux.HandleFunc("/", root)
	mux.HandleFunc("/movie", app.getOneMovieJSON)
	mux.HandleFunc("/movies", app.getMoviesJSON)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./public/favicon.ico")
	})
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)
  	n.Run(":" + app.config.Web.Port)

	app.server = n
}
