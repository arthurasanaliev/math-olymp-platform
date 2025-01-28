package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"net/http"
)

// Blogs is the home page handler
func (m *Repository) Blogs(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
        "Title": "Blogs",
	}
	render.RenderTemplate(w, "blogs.html", data)
}

