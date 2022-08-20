package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(a.loadSession)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {

		// a.session.Put(r.Context(), "test", "lsdjf")

		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Println(err)
		}
	})

	return mux
}
