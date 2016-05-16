package main

import (
	"strconv"

	"github.com/dinever/golf"
)

func (app *application) root(ctx *golf.Context) {
	data := map[string]interface{}{
    	"Title": "Hello World",
	}
	ctx.Loader("template").Render("index.html", data)
}

func (app *application) getOneMovieJSON(ctx *golf.Context) {
	movies := app.getMovies(1, 0)
	ctx.JSON(movies)
}

func (app *application) getMoviesJSON(ctx *golf.Context) {
	limitStr := ctx.Param("limit")
	offsetStr := ctx.Param("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	data := app.getMovies(limit, offset)
	ctx.JSON(data)
}
