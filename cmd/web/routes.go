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
	mux.Get("/blogs", handlers.Repo.Blogs)
	mux.Get("/problems", handlers.Repo.Problems)
	mux.Get("/problems/{id}", handlers.Repo.Problem)
	mux.Get("/login", handlers.Repo.Login)
	mux.Get("/signup", handlers.Repo.Signup)
	mux.Get("/about", handlers.Repo.About)

	mux.Post("/login", handlers.Repo.Login)
	mux.Post("/signup", handlers.Repo.Signup)
	mux.Post("/problems/{id}", handlers.Repo.Problem)

	return mux
}
