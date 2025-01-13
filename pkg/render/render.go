package render

import (
	"html/template"
	"net/http"
)

// RenderTemplate renders a template with the given data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates, err := template.ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/"+tmpl,
	)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	err = templates.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
