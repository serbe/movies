package main

import (
	"strconv"
		"encoding/json"

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
	data, err := json.Marshal(movies)
	if err != nil {
		panic(err)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Send(data)

}

func (app *application) getMoviesJSON(ctx *golf.Context) {
	limitStr := ctx.Param("limit")
	offsetStr := ctx.Param("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	movies := app.getMovies(limit, offset)
	data, err := json.Marshal(movies)
	if err != nil {
		panic(err)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Send(data)
}
