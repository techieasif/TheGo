package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

func NoSurfCSRF(next http.Handler) http.Handler {
	noSurfHandler := nosurf.New(next)
	noSurfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	return noSurfHandler
}
