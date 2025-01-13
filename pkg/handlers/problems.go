package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"net/http"
)

// Problems is the problems page handler
func (m *Repository) Problems(w http.ResponseWriter, r *http.Request) {
	problems, err := m.app.DB.GetAllProblems()
	if err != nil {
		http.Error(w, "Error fetching problems: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.RenderTemplate(w, "problems.html", problems)
}
