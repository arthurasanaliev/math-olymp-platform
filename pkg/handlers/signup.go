package handlers

import "net/http"

// Signup is the signup page handler
func (m *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "signup.html", nil)
		if err != nil {
			http.Error(w, "Unable to load the signup form", http.StatusInternalServerError)
		}
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
			err := tmpl.ExecuteTemplate(w, "signup.html", map[string]string{
				"Error": "This username is already taken",
			})
			if err != nil {
				http.Error(w, "Unable to load the signup form", http.StatusInternalServerError)
			}
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
