package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mcsymiv/go-stripe/internal/web/config"
	"github.com/mcsymiv/go-stripe/internal/web/handlers"
)

// routes
// from chi documentation:
// NewMux returns a newly initialized Mux object that implements the Router interface
// Mux objects implements Handler ServeHTTP method
func routes(a *config.Application) http.Handler {
	var mux *chi.Mux = chi.NewMux()
	mux.Get("/virtual-terminal", handlers.Repo.VirtualTerminal)
	mux.Post("/payment-succeeded", handlers.Repo.PaymentSucceeded)

	fs := http.FileServer(http.Dir("/static/*"))
	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	return mux
}
