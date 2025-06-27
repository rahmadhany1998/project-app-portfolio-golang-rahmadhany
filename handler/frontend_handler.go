package handler

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"project-app-portfolio-golang-rahmadhany/model"
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

func (h *FrontendHandler) ShowPortfolioDetail(w http.ResponseWriter, r *http.Request) {
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

func (h *FrontendHandler) ShowContactForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/contact.html",
	))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func (h *FrontendHandler) SubmitContactForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	contact := model.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	// Validasi
	if contact.Name == "" || contact.Email == "" || contact.Subject == "" || contact.Message == "" {
		http.Error(w, "Semua field wajib diisi", http.StatusBadRequest)
		return
	}

	err = h.apiService.SubmitContact(contact)
	if err != nil {
		http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contact", http.StatusSeeOther)
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

func (h *FrontendHandler) ShowAddPortfolioForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/portfolio_add.html",
	))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func (h *FrontendHandler) SubmitPortfolio(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	shortDesc := r.FormValue("short_description")
	client := r.FormValue("client")
	website := r.FormValue("website")
	longDesc := r.FormValue("long_description")

	if title == "" || shortDesc == "" || client == "" || website == "" || longDesc == "" {
		http.Error(w, "Semua field harus diisi", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image_file")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := handler.Filename
	image := "/static/img/portfolio/" + filename
	savePath := "web/static/img/portfolio/" + filename
	out, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	portfolio := model.Portfolio{
		Title:            title,
		Image:            image,
		ShortDescription: shortDesc,
		Client:           client,
		Website:          website,
		LongDescription:  longDesc,
	}

	err = h.apiService.AddPortfolio(portfolio)
	if err != nil {
		http.Error(w, "Failed to save portfolio", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
}
