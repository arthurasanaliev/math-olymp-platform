package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/config"
)

// Repository stores data for handlers package
type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// SetRepo sets Repository instance
func SetRepo(r *Repository) {
	Repo = r
}
