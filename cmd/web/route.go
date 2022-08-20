package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (a *application) route() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Raihan"))
	})
	return mux
}
