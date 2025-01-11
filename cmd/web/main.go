package main

import (
	"context"
	"github.com/arthurasanaliev/math-olymp-platform/pkg/config"
	"github.com/arthurasanaliev/math-olymp-platform/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig

// main is the entry point of the program
func main() {
	conn, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	app = config.AppConfig{
		InProduction: false,
		DB:           conn,
	}

	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("Running app on port" + portNumber)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
