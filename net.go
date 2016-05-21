package main

import (
	"net/http"
	"time"
)

func (app *application) initServer() {
	var ServeMux = http.NewServeMux()

	ServeMux.HandleFunc("/", app.root)
	ServeMux.HandleFunc("/movie", app.getOneMovieJSON)
	ServeMux.HandleFunc("/movies", app.getMoviesJSON)
	ServeMux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./public/favicon.ico")
	})
	ServeMux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	s := &http.Server{
		Addr:           ":" + app.config.Web.Port,
		Handler:        ServeMux,         // handler to invoke, http.DefaultServeMux if nil
		ReadTimeout:    10 * time.Second, // maximum duration before timing out read of the request
		WriteTimeout:   10 * time.Second, // maximum duration before timing out write of the response
		MaxHeaderBytes: 1 << 20,          // maximum size of request headers, 1048576 bytes
	}
	s.ListenAndServe()
	app.server = s
}
