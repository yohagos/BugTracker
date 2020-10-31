package utils

import (
	"context"
	"html/template"
	"net/http"
)

var ctx = context.TODO()
var templates *template.Template

// LoadTemplate func
func LoadTemplate(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

// ExecuteTemplate func
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}
