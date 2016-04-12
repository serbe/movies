package main

import (
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
