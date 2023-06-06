package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/techieasif/TheGo/bookings/internal/config"
	"github.com/techieasif/TheGo/bookings/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

var pathToTemplates = "./templates"

func NewTemplates(a *config.AppConfig) {
	app = a
}

var tc = make(map[string]*template.Template)

func AddDefaultTemplateData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func ParseTemplateV2(w http.ResponseWriter, t string, td *models.TemplateData, r *http.Request) error {
	//create template cache

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCacheV2()
	}

	// Get requested template
	tmpl, k := tc[t]
	if !k {
		return errors.New("could not find template in cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultTemplateData(td, r)

	err := tmpl.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	return nil
}
func CreateTemplateCacheV2() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layoutMaches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			return myCache, err
		}
		if len(layoutMaches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}
