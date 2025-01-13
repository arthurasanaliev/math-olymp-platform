package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// Problem is the problem page handler
func (m *Repository) Problem(w http.ResponseWriter, r *http.Request) {
	vars := chi.URLParam(r, "id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		return
	}

	problem, err := m.app.DB.GetProblemByID(id)
	if err != nil {
		http.Error(w, "Problem not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		userAnswer, err := strconv.Atoi(r.FormValue("answer"))
		if err != nil {
			http.Error(w, "Invalid answer format", http.StatusBadRequest)
			return
		}

		correctAnswer := problem.Answer

		message := "Incorrect"
		if userAnswer == correctAnswer {
			message = "Correct"
		}

		data := map[string]interface{}{
			"Title":   "Problem",
			"Problem": problem,
			"Message": message,
		}
		render.RenderTemplate(w, "problem.html", data)
		return
	}

	data := map[string]interface{}{
		"Title":   "Problem",
		"Problem": problem,
	}
	render.RenderTemplate(w, "problem.html", data)
}
