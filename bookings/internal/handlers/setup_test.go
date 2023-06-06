package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/techieasif/TheGo/bookings/internal/config"
	"github.com/techieasif/TheGo/bookings/internal/render"
	"github.com/techieasif/TheGo/bookings/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"

var functions = template.FuncMap{}

func getRoutes() http.Handler {
	//What I am going to put in session
	gob.Register(models.Reservation{})

	app.InProduction = false
	session = scs.New()

	//24 hour session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCacheV2()

	if err != nil {
		log.Fatal(err.Error())
	}
	app.TemplateCache = tc
	app.UseCache = true
	repo := NewRepo(&app)
	NewHandlers(repo)
	render.NewTemplates(&app)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.DefaultLogger)
	//mux.Use(NoSurfCSRF)
	mux.Use(SessionLoad)
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.PostAvailabilityJson)

	filerServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", filerServer))
	return mux

}

//NoSurfCSRF add CSRF protection
func NoSurfCSRF(next http.Handler) http.Handler {
	noSurfHandler := nosurf.New(next)
	noSurfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	return noSurfHandler
}

//SessionLoad loads the session on each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCacheV2() (map[string]*template.Template, error) {
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
