package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func (a *application) noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   a.config.inProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func (a *application) sessionLoad(next http.Handler) http.Handler {
	return a.session.LoadAndSave(next)
}
