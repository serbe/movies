package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		"title": "Movie RSS",
	})
}

func (app *application) getOneMovieJSON(c *gin.Context) {
	movies := app.getMovies(1, 0)
	c.JSON(200, movies)
}

func (app *application) getMoviesJSON(c *gin.Context) {
	limitStr, _ := c.GetQuery("limit")
	offsetStr, _ := c.GetQuery("offset")
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	data := app.getMovies(limit, offset)
	c.JSON(200, data)
}
