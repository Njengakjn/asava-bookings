package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(a *application) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(a.noSurf)
	mux.Use(a.sessionLoad)

	mux.Get("/v1/", a.homePageHandlers)
	mux.Get("/v1/about", a.aboutPageHandlers)

	return mux
}
