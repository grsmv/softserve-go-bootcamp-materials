package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(template.ParseFiles(filepath.Join(TemplatesPath, "edit.tmpl"), filepath.Join(TemplatesPath, "view.tmpl")))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".tmpl", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
