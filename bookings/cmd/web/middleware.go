package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

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
