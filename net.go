package main

import (
	"github.com/dinever/golf"
)

func (app *application) initServer() {
	g := golf.New()
	g.View.SetTemplateLoader("template", "templates/")
  	g.Get("/", app.root)
	g.Get("/movie", app.getOneMovieJSON)
	g.Get("/movies", app.getMoviesJSON)
	// mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
	// 	http.ServeFile(w, req, "./public/favicon.ico")
	// })
	g.Static("/public/", "public")

  	g.Run(":" + app.config.Web.Port)

	app.server = g
}
