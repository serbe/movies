package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		"title": "Movie RSS",
	})
}

func react(c *gin.Context) {
	c.HTML(200, "react", gin.H{
		"title": "React Tutorial",
	})
}

func (app *application) getOneMovieJSON(c *gin.Context) {
	movie := app.getMovie()
	c.JSON(200, movie)
}

func (app *application) getMoviesJSON(c *gin.Context) {
	movies := app.getMovies()
	fmt.Printf("%+v",movies[0])
	c.JSON(200, movies)
}