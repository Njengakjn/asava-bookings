package main

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func renderTemplate(w http.ResponseWriter, tc map[string]*template.Template, td *templateData, name string) error {
	tmpl, ok := tc[name]
	if !ok {
		return errors.New("missing template")
	}

	buf := new(bytes.Buffer)

	err := tmpl.Execute(buf, td)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
