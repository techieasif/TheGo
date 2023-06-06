package models

import "github.com/techieasif/TheGo/bookings/internal/forms"

//TemplateData is the data passed to every page and is available to use in templates.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
