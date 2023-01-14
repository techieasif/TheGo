package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/contact-us", contactUS)
	_ = http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	renderTemplate(w, "home.page.tmpl")

}

func contactUS(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "You can't contact us says chhota don")
	if err != nil {
		fmt.Println("Some error", err.Error())
	}

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occurred while executing template:", err)
		return
	}

}
