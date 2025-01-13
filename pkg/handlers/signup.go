package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/render"
	"net/http"
)

// Signup is the signup page handler
func (m *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := map[string]string{
			"Title": "Sign Up",
		}
		render.RenderTemplate(w, "signup.html", data)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		isValidUser, err := m.app.DB.CheckUserExists(username)
		if err != nil {
			http.Error(w, "Error verifying user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if isValidUser {
			data := map[string]string{
				"Error": "This username is already taken",
			}
			render.RenderTemplate(w, "signup.html", data)
		} else {
			err := m.app.DB.InsertUser(username, password)
			if err != nil {
				http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		return
	}
}
