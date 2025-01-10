package main

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/config"
	"github.com/arthurasanaliev/math-olymp-platform/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

// main is the entry point of the program
func main() {
	app.InProduction = false

	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("Running app on port" + portNumber)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
