package handlers

import (
	"github.com/techieasif/TheGo/bookings/models"
	"github.com/techieasif/TheGo/bookings/pkg/config"
	"github.com/techieasif/TheGo/bookings/pkg/render"
	"net/http"
)

//Repo is the repository used by handlers.
var Repo *Repository

// Repository is the repository type.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplateV2(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["name"] = "asif"
	render.ParseTemplateV2(w, "about.page.tmpl", &models.TemplateData{StringMap: strMap})
}
