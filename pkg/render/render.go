package render

import (
	"bytes"
	"github.com/danielleedottech/bookings/config"
	"github.com/danielleedottech/bookings/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}
	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("template not found: ", tmpl)
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, AddDefaultData(td))

	if err != nil {
		log.Fatal(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.gohtml")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			tsWithLayouts, err := ts.ParseGlob("./templates/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}

			myCache[name] = tsWithLayouts
		} else {
			myCache[name] = ts
		}
	}

	return myCache, nil
}
