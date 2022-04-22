package main

import (
	. "github.com/RafaelRCamargo/go-svg-api/server"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	Server(port)
}
