package main

import (
	"log"
	"net/http"
	"newsapi/internal/router"
)

func main() {

	r := router.New()

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Failed to to Start the Server", err)
	}
}
