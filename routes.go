package main

import (
	"net/http"
	"strconv"
)

func (app *application) root(w http.ResponseWriter, req *http.Request) {
	app.render(w, ctx{title: "Movie RSS"}, "index") 
}

func (app *application) getOneMovieJSON(w http.ResponseWriter, req *http.Request) error {
	movies := app.getMovies(1, 0)
	return c.JSON(200, movies)
}

func (app *application) getMoviesJSON(w http.ResponseWriter, req *http.Request) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	data := app.getMovies(limit, offset)
	return c.JSON(200, data)
}
