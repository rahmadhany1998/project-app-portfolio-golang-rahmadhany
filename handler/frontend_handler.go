package handler

import (
	"html/template"
	"net/http"

	"project-app-portfolio-golang-rahmadhany/service"
)

type FrontendHandler struct {
	tmpl       *template.Template
	apiService service.ApiService
}

func NewFrontendHandler(apiService service.ApiService) *FrontendHandler {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/index.html",
	))

	return &FrontendHandler{
		tmpl:       tmpl,
		apiService: apiService,
	}
}

func (h *FrontendHandler) ShowHome(w http.ResponseWriter, r *http.Request) {
	user, err := h.apiService.GetProfile()
	if err != nil {
		http.Error(w, "failed to load user profile", http.StatusInternalServerError)
		return
	}

	h.tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Name": user.Name,
		"Job":  user.Job,
	})
}

func (h *FrontendHandler) ShowContactForm(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "layout", nil)
}

func (h *FrontendHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	// email := r.FormValue("email")
	// message := r.FormValue("message")

	w.Write([]byte("Thanks, " + name + ". Your message has been received."))
}

func (h *FrontendHandler) ShowAbout(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "layout", nil)
}
