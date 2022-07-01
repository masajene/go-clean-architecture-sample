package main

import (
	_ "embed"
	"go_api_boilerplate/infra/router"
	"log"
)

const PORT = "3000"

func main() {

	// init router
	r := router.NewRouter()

	// start server
	err := r.API.Run("127.0.0.1:" + PORT)
	if err != nil {
		log.Fatalf("faild to run: %v", err)
	}
}
