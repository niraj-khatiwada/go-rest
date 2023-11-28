package utils

import (
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

func GetRouterMiddlewares() []func(http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{chiMiddleware.RealIP, chiMiddleware.StripSlashes, chiMiddleware.RedirectSlashes}
	if os.Getenv("IS_DEBUG_MODE") == "true" {
		middlewares = append(middlewares, chiMiddleware.Logger)
	}
	return middlewares
}
