package main

import (
	"log"
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

	fileServer(mux)

	mux = endPoints(mux)
	return mux
}

// fileServer serves the files in the /static/ folder
func fileServer(mux *chi.Mux) {
	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fs))
	log.Println(fs)
}

// endPoints returns routes of the app
func endPoints(mux *chi.Mux) *chi.Mux {
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/room-family-first", handlers.Repo.RoomFamilyFirst)
	mux.Get("/room-romantic-getaway", handlers.Repo.RoomRomanticGetaway)
	mux.Get("/room-royal-garden", handlers.Repo.RoomRoyalGarden)
	mux.Get("/room-your-lady", handlers.Repo.RoomYourLady)
	mux.Get("/room-your-majesty", handlers.Repo.RoomYourMajesty)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	return mux
}
