package api

import (
	"github.com/go-chi/chi/v5"
	api "github.com/niraj-khatiwada/go-rest/api/v1/auth"
)

func V1(r chi.Router) {
	r.Route("/v1", func(router chi.Router) {
		api.Auth(router)
	})
}
