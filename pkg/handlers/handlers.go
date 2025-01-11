package handlers

import (
	"github.com/arthurasanaliev/math-olymp-platform/pkg/config"
	"html/template"
	"net/http"
)

// Repository stores data for handlers package
type Repository struct {
	app *config.AppConfig
}

var (
	Repo *Repository
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
)

// NewRepo creates a new Repository instance
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

// SetRepo sets Repository instance
func SetRepo(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "Unable to load the home page", http.StatusInternalServerError)
	}
}

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
