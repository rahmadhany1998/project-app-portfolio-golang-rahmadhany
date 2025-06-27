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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		h.Frontend.ShowHome(w, r)
	})

	r.Get("/portfolio", func(w http.ResponseWriter, r *http.Request) {
		h.Frontend.ShowPortfolio(w, r)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/profile", h.Api.GetProfile)
		r.Get("/portfolios", h.Api.GetPortfolios)
	})

	return r
}
