package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/niraj-khatiwada/go-rest/api/v1"
	"github.com/niraj-khatiwada/go-rest/middlewares"
)

func Api(r chi.Router) {
	r.Route("/api", func(router chi.Router) {
		header := middlewares.Header{}
		router.Use(header.WithSetHeader("Content-Type", "application/json"))
		api.V1(router)
	})
}
