package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/techieasif/TheGo/bookings/internal/config"
	"github.com/techieasif/TheGo/bookings/internal/handlers"
	"github.com/techieasif/TheGo/bookings/internal/render"
	"github.com/techieasif/TheGo/bookings/models"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
}

func run() error {
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

	tc, err := render.CreateTemplateCacheV2()

	if err != nil {
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	return nil
}
