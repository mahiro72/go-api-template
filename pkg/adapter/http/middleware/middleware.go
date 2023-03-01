package middleware

import (
	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func SetCommonMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

}
