package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/niraj-khatiwada/go-rest/utils"
	"net/http"
)

func Routes(r chi.Router) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("<h1>404: Page not found</h1>")); err != nil {
			utils.CatchRuntimeErrors(err)
			http.Error(w, "error", 500)
			return
		}
	})
}
