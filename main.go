package main

import (
	"GoRestApi/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", router))
}
