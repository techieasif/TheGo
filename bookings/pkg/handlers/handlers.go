package handlers

import (
	"github.com/techieasif/TheGo/bookings/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "about.page.tmpl")
}
