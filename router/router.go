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
	r.Get("/portfolio", h.Frontend.ShowPortfolio)
	r.Get("/portfolio/{id}", h.Frontend.ShowPortfolioDetail)
	r.Get("/portfolio/add", h.Frontend.ShowAddPortfolioForm)
	r.Post("/portfolio/add", h.Frontend.SubmitPortfolio)
	r.Get("/contact", h.Frontend.ShowContactForm)
	r.Post("/contact", h.Frontend.SubmitContactForm)
	r.Get("/about", h.Frontend.ShowAbout)
	r.Get("/experience", h.Frontend.ShowExperience)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/profile", h.Api.GetProfile)
		r.Get("/portfolios", h.Api.GetPortfolios)
		r.Get("/portfolios/{id}", h.Api.GetPortfolioDetail)
		r.Post("/portfolios/add", h.Api.CreatePortfolio)
		r.Post("/contact", h.Api.SubmitContact)
		r.Get("/about", h.Api.GetAbout)
		r.Get("/experience", h.Api.GetAllExperiences)
	})

	return r
}
