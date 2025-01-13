package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"net/http"
)

// Problemset is the problemset page handler
func (m *Repository) Problemset(w http.ResponseWriter, r *http.Request) {
	problems, err := m.app.DB.GetAllProblems()
	if err != nil {
		http.Error(w, "Error fetching problems: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.RenderTemplate(w, "problemset.html", problems)
}
