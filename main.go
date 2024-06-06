package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define the port the server will listen on
const port = ":8080"

// Home handler function
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

// Info handler function
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info page")
}

func main() {
	// Log that the server is starting
	log.Println("Starting the HTTP server on port", port)

	// Register handler functions for specific routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	// Start the server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
