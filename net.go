package main

import (
	"net/http"
)

func (app *application) initServer() {
  	// http.Handle("/", http.FileServer(http.Dir("./templates/root.html")))
	http.HandleFunc("/movie", app.getOneMovieJSON)
	http.HandleFunc("/movies", app.getMoviesJSON)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./public/favicon.ico")
	})
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

  	http.ListenAndServe(":" + app.config.Web.Port, nil)
}
