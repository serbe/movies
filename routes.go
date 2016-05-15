package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func root(c echo.Context) error {
	return c.Render(200, "index", gin.H{
		"title": "Movie RSS",
	})
}

func (app *application) getOneMovieJSON(c echo.Context) error {
	movies := app.getMovies(1, 0)
	return c.JSON(200, movies)
}

func (app *application) getMoviesJSON(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	data := app.getMovies(limit, offset)
	return c.JSON(200, data)
}
