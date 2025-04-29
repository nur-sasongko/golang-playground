package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoHandler)

	log.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
