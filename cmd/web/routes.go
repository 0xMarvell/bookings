package main

import (
	"net/http"

	"github.com/Marvellous-Chimaraoke/bookings/pkg/config"
	"github.com/Marvellous-Chimaraoke/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

// routes handles URL routing for all handlers created
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Add some middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
