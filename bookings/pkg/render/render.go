package render

import (
	"bytes"
	"fmt"
	"github.com/techieasif/TheGo/bookings/models"
	"github.com/techieasif/TheGo/bookings/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// ParseTemplate is used to parse go templates
func ParseTemplateT(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the templates:", err)
	}
}

var tc = make(map[string]*template.Template)

//ParseTemplateV1 v1
func ParseTemplateV1(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//Checking if template is already cached
	_, inMap := tc[t]

	if !inMap {
		// template is not cached, create template cache.
		log.Println("Creating template and adding to cache")
		err := createTemplateCache(t)
		if err != nil {
			log.Println("Error while creating template")
		}

	} else {
		// already in cache, get it and cache.

		log.Println("From cached template", t)

	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

//createTemplateCache v1
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
	}

	fmt.Println("Templates: ", templates)
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}

///V2 code

func AddDefaultTemplateData(td *models.TemplateData) *models.TemplateData {
	return td
}

func ParseTemplateV2(w http.ResponseWriter, t string, td *models.TemplateData) {
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
		log.Fatal("Could not find template in cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultTemplateData(td)

	err := tmpl.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}
func CreateTemplateCacheV2() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layoutMaches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}
		if len(layoutMaches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}
