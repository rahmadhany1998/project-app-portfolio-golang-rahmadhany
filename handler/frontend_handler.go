package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"project-app-portfolio-golang-rahmadhany/service"

	"github.com/go-chi/chi"
)

type FrontendHandler struct {
	apiService service.ApiService
}

func NewFrontendHandler(apiService service.ApiService) *FrontendHandler {
	return &FrontendHandler{
		apiService: apiService,
	}
}

func (h *FrontendHandler) ShowHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/index.html",
	))

	user, err := h.apiService.GetProfile()
	if err != nil {
		http.Error(w, "failed to load user profile", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Name": user.Name,
		"Job":  user.Job,
	})
}

func (h *FrontendHandler) ShowPortfolio(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/portfolio.html",
	))

	items, err := h.apiService.GetPortfolios()
	if err != nil {
		http.Error(w, "failed to load portfolio", http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Portfolios": items,
	})
}

func (h *FrontendHandler) ShowPortfolioDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"web/templates/layout.html",
			"web/templates/header.html",
			"web/templates/footer.html",
			"web/templates/portfolio_detail.html",
		))

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid portfolio id", http.StatusBadRequest)
			return
		}

		item, err := h.apiService.GetPortfolioByID(id)
		if err != nil {
			http.Error(w, "portfolio not found", http.StatusNotFound)
			return
		}

		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
			"Portfolio": item,
		})
	}
}

func (h *FrontendHandler) ShowContactForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/contact.html",
	))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func (h *FrontendHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	w.Write([]byte("Thanks, " + name + ". Your message has been received."))
}

func (h *FrontendHandler) ShowAbout(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/about.html",
	))
	tmpl.ExecuteTemplate(w, "layout", nil)
}
