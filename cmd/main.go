package main

import (
	"log"
	"net/http"
	"studies-1/internals/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
