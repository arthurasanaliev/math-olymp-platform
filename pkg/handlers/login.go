package handlers

import "net/http"

// Login is the login page handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, "Unable to load the login form", http.StatusInternalServerError)
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

		isValidPass := false
		if isValidUser {
			isValidPass, err = m.app.DB.CheckUserPassword(username, password)
			if err != nil {
				http.Error(w, "Error verifying user: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if isValidUser && isValidPass {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			err := tmpl.ExecuteTemplate(w, "login.html", map[string]string{
				"Error": "Invalid username or password",
			})
			if err != nil {
				http.Error(w, "Unable to load the login form", http.StatusInternalServerError)
			}
		}
		return
	}
}
