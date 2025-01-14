package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"net/http"
)

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "About",
	}
	render.RenderTemplate(w, "about.html", data)
}
