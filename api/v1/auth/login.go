package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Login(r chi.Router) {
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {

		person := map[string]string{"name": "Niraj", "age": "26"}
		jsan, _ := json.Marshal(person)
		if _, err := w.Write(jsan); err != nil {

		}
	})
}
