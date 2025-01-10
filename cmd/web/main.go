package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point of the program
func main() {
	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

    log.Println("Running app on port"+portNumber)

    if err := server.ListenAndServe(); err != nil {
        log.Fatal("Error starting server:", err)
    }
}
