package router

import (
	"project-app-portfolio-golang-rahmadhany/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.User.GetAll)
		r.Post("/", h.User.Create)
		r.Get("/{id}", h.User.GetByID)
		r.Put("/{id}", h.User.Update)
		r.Delete("/{id}", h.User.Delete)
	})

	return r
}
