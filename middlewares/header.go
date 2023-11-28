package middlewares

import "net/http"

type Header struct {
}

/*
WithSetHeader : Set a header to whole route.
*/
func (header *Header) WithSetHeader(key string, value string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			next.ServeHTTP(w, r)
		})
	}
}
