package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Server() {
	println("Starting server...")

	err := godotenv.Load(".env")

	if err != nil {
		println("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<div>aaa</div>",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	r.Run(os.Getenv("PORT"))

	println("Listing for requests at http://localhost/" + os.Getenv("PORT"))
}
