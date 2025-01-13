package main

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes defines routing
func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(noSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/login", handlers.Repo.Login)
	mux.Get("/signup", handlers.Repo.Signup)
	mux.Get("/problems", handlers.Repo.Problems)

	mux.Post("/login", handlers.Repo.Login)
	mux.Post("/signup", handlers.Repo.Signup)

	return mux
}
