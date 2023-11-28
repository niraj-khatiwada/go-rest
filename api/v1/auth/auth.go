package api

import "github.com/go-chi/chi/v5"

func Auth(r chi.Router) {
	r.Route("/auth", func(router chi.Router) {
		Login(router)
	})
}
