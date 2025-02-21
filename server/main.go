package main

import (
	"fmt"
	"net/http"
	"log"
)

// Handler function for the root endpoint
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The server is Running!")
}

// Handler para la ruta "/print"
func printHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Called path: /print")
	fmt.Fprintf(w, "Print handler executed!")
}

func main() {
	// Register handler for the root URL path
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/print", printHandler)

	// Start the server on port 8080
	fmt.Println("Server running on http://localhost:8080/print")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}