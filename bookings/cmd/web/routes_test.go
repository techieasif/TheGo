package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/techieasif/TheGo/bookings/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	//do nothing
	default:
		fmt.Println(fmt.Sprintf("mux is not type of chi mux, instead it is off %T", v))
	}
}
