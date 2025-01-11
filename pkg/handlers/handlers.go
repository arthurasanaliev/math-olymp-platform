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

// Home is the home page handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "Unable to load the home page", http.StatusInternalServerError)
	}
}

// Login is the login page handler function
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

		if isValidUser {
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
