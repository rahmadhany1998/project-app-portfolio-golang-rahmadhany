package router

import (
	"net/http"
	"project-app-portfolio-golang-rahmadhany/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	r.Get("/", h.Frontend.ShowHome)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/profile", h.Api.GetProfile)
	})

	return r
}
