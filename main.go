package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<div>aaa</div>",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./favicon.ico")
	})

	r.Run(":8000")

	log.Println("Listing for requests at http://localhost:8000/")
}
