package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/snipersap/bookings/pkg/config"
	"github.com/snipersap/bookings/pkg/handlers"
)

// routes creates a new router, registers the middlewares and loads the session
func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(LogEachApiHit)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux = endPoints(mux)
	return mux
}

// endPoints returns routes of the app
func endPoints(mux *chi.Mux) *chi.Mux {
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
