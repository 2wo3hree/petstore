package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggo/http-swagger"
	"repo/internal/handler"
)

func SetupRouter(contr *handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", contr.Create)
		r.Get("/{id}", contr.GetByID)
		r.Put("/{id}", contr.Update)
		r.Delete("/{id}", contr.Delete)
		r.Get("/", contr.List)
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
