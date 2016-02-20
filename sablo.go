package sablo

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates map[string]*template.Template

// LoadTemplates creates template cache
func LoadTemplates(layoutPattern string, pagePattern string) error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layouts, err := filepath.Glob(layoutPattern)
	if err != nil {
		return nil
	}

	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return nil
	}

	for _, page := range pages {
		files := append(layouts, page)
		templates[filepath.Base(page)] = template.Must(template.ParseFiles(files...))
	}
	return nil
}

// RenderTemplate is a wrapper around template.ExecuteTemplate.
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		panic("The template " + name + " does not exist.")
	}
	tmpl.ExecuteTemplate(w, "base", data)
}
