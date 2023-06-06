package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/techieasif/TheGo/bookings/internal/config"
	"github.com/techieasif/TheGo/bookings/internal/forms"
	"github.com/techieasif/TheGo/bookings/internal/render"
	"github.com/techieasif/TheGo/bookings/models"
	"log"
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
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.ParseTemplateV2(w, "home.page.tmpl", &models.TemplateData{}, r)
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["name"] = "asif"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	strMap["remote_ip"] = remoteIP
	render.ParseTemplateV2(w, "about.page.tmpl", &models.TemplateData{StringMap: strMap}, r)
}

// Reservation is the handler for the Reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.ParseTemplateV2(w, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, r)
}

// PostReservation handles the posting of forms
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error while parsing the form", err.Error())
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	//form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})

		data["reservation"] = reservation

		render.ParseTemplateV2(w, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		}, r)

		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Println("Error occurred while fetching data from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})

	data["reservation"] = reservation

	render.ParseTemplateV2(w, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	}, r)
}

// Generals is the handler for the Generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplateV2(w, "generals.page.tmpl", &models.TemplateData{}, r)
}

// Majors is the handler for the Majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplateV2(w, "majors.page.tmpl", &models.TemplateData{}, r)
}

// Availability is the handler for the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplateV2(w, "search-availability.page.tmpl", &models.TemplateData{}, r)
}

// PostAvailability is the handler for the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start:", r.Form.Get("start"), "end:", r.Form.Get("end"))
	_, err := w.Write([]byte("Post search-availability"))
	if err != nil {
		return
	}
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostAvailabilityJson(w http.ResponseWriter, r *http.Request) {

	resp := &jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, e := json.Marshal(resp)

	if e != nil {
		fmt.Println("error while sending json", e.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplateV2(w, "contact.page.tmpl", &models.TemplateData{}, r)
}
