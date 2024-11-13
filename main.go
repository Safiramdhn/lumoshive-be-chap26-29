package main

import (
	"golang-beginner-chap28/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.NewRouter()

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
		log.Fatalf("Error starting server: %v\n", err.Error())
	}

	http.ListenAndServe(":8080", r)
}
