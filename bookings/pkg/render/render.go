package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// ParseTemplate is used to parse go templates
func ParseTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the templates:", err)
	}
}
