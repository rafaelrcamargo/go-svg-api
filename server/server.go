package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Server() {
	println("Starting server...")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<div>aaa</div>",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	r.Run(":80")

	println("Listing for requests at http://localhost/")
}
