package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/masajene/go-clean-architecture-sample/infrastructure"
)

func main() {
	loadEnv()

	port := ":" + os.Getenv("PORT")
	infrastructure.Router.Run(port)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
