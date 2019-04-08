package sablo

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templateCache map[string]*template.Template
var pageCache map[string]*template.Template

// LoadTemplates creates template cache
func LoadTemplates(layoutPattern string, pagePattern string) error {
	if templateCache == nil {
		templateCache = make(map[string]*template.Template)
	}

	layouts, err := filepath.Glob(layoutPattern)
	if err != nil {
		return err
	}

	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return err
	}

	for _, page := range pages {
		files := append(layouts, page)
		templateCache[filepath.Base(page)] = template.Must(template.ParseFiles(files...))
	}
	return nil
}

// LoadPages loads page files without a template
func LoadPages(pagePattern string) error {
	if pageCache == nil {
		pageCache = make(map[string]*template.Template)
	}

	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return err
	}
	for _, page := range pages {
		pageCache[filepath.Base(page)] = template.Must(template.ParseFiles(page))
	}
	return nil
}

// RenderPage renders single page
func RenderPage(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := pageCache[name]
	if !ok {
		panic("The template " + name + " does not exist.")
	}
	return tmpl.Execute(w, data)
}

// RenderTemplate is a wrapper around template.ExecuteTemplate.
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templateCache[name]
	if !ok {
		panic("The template " + name + " does not exist.")
	}
	return tmpl.ExecuteTemplate(w, "base", data)
}
