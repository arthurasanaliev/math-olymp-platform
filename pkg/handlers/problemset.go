package handlers

import "net/http"

// Problemset is the problemset page handler
func (m *Repository) Problemset(w http.ResponseWriter, r *http.Request) {
	problems, err := m.app.DB.GetAllProblems()
	if err != nil {
		http.Error(w, "Error fetching problems: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "problemset.html", problems)
	if err != nil {
		http.Error(w, "Unable to load the problemset page", http.StatusInternalServerError)
		return
	}
}
