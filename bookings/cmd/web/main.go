package main

import (
	"fmt"
	"github.com/techieasif/TheGo/bookings/pkg/config"
	"github.com/techieasif/TheGo/bookings/pkg/handlers"
	"github.com/techieasif/TheGo/bookings/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCacheV2()

	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	//_ = http.ListenAndServe(portNumber, nil)
}
