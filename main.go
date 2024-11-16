package main

import (
	"golang-beginner-chap28/config"
	"golang-beginner-chap28/routers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Initialize configuration and router
	configViper := config.InitViper()
	r := routers.NewRouter()

	// Determine server port
	port := configViper.GetInt("PORT")
	if port == 0 {
		port = 8081 // Default port
	}
	portStr := ":" + strconv.Itoa(port)

	// Log server startup
	log.Printf("Server started on port %d", port)

	// Start the server
	if err := http.ListenAndServe(portStr, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
