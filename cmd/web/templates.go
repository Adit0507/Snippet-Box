package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Adit0507/Snippet-Box/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	// field for holding a slice of snippets
	Snippets []*models.Snippet
	Form     any
	Flash    string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2024 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Map to act as the cache
	cache := map[string]*template.Template{}

	// filePath.Glob()is used to get a slice of all filepaths that match the
	// pattern "./ui/html/pages/*.tmpl", like  [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// looping through filepaths one by one
	for _, page := range pages {
		// extraccting file name(eg: home.tmpl) & assigning
		// it to name variable
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
