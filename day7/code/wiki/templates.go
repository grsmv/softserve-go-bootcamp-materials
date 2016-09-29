package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type render interface {
	renderTemplate(w http.ResponseWriter, tmpl string, p *Page)
}

type compiledTemplates struct {
	templates *template.Template
}

func NewCTemplates(templatesPath string) *compiledTemplates {
	return &compiledTemplates{
		templates: template.Must(template.ParseFiles(filepath.Join(templatesPath, "edit.tmpl"), filepath.Join(templatesPath, "view.tmpl"))),
	}
}

func (c *compiledTemplates) renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := c.templates.ExecuteTemplate(w, tmpl+".tmpl", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type DirectTemplate struct {
	path string
}

func (dt DirectTemplate) renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(filepath.Join(dt.path, tmpl+".tmpl"))
	t.Execute(w, p)
}

var templates render
