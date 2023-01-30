package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

// routes
// from chi documentation:
// NewMux returns a newly initialized Mux object that implements the Router interface
// Mux objects implements Handler ServeHTTP method
func (a *app) routes() http.Handler {
	var mux *chi.Mux = chi.NewMux()
	mux.Get("/virtual-terminal", a.VirtualTerminal)
	return mux
}
