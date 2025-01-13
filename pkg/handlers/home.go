package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"net/http"
)

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Home - Math Olymp Platform",
	}
	render.RenderTemplate(w, "home.html", data)
}
