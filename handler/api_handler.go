package handler

import (
	"io"
	"net/http"
	"os"
	"project-app-portfolio-golang-rahmadhany/model"
	"project-app-portfolio-golang-rahmadhany/service"
	"project-app-portfolio-golang-rahmadhany/util"
	"strconv"

	"github.com/go-chi/chi"
)

type ApiHandler struct {
	service service.ApiService
}

func NewApiHandler(s service.ApiService) *ApiHandler {
	return &ApiHandler{s}
}

func (h *ApiHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.GetProfile()
	if err != nil {
		util.WriteError(w, "failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	util.WriteSuccess(w, "user profile fetched", map[string]interface{}{
		"name": user.Name,
		"job":  user.Job,
	})
}

func (h *ApiHandler) GetPortfolios(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetPortfolios()
	if err != nil {
		util.WriteError(w, "failed to fetch portfolios", http.StatusInternalServerError)
		return
	}

	util.WriteSuccess(w, "portfolio list", items)
}

func (h *ApiHandler) GetPortfolioDetail(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.WriteError(w, "invalid portfolio id", http.StatusBadRequest)
		return
	}

	portfolio, err := h.service.GetPortfolioByID(id)
	if err != nil {
		util.WriteError(w, "portfolio not found", http.StatusNotFound)
		return
	}

	util.WriteSuccess(w, "portfolio detail fetched", portfolio)
}

func (h *ApiHandler) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		util.WriteError(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	shortDesc := r.FormValue("short_description")
	client := r.FormValue("client")
	website := r.FormValue("website")
	longDesc := r.FormValue("long_description")

	if title == "" || shortDesc == "" || client == "" || website == "" || longDesc == "" {
		util.WriteError(w, "Semua field harus diisi", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image_file")
	if err != nil {
		util.WriteError(w, "File image wajib diunggah", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := handler.Filename
	image := "/static/img/portfolio/" + filename
	savePath := "web/static/img/portfolio/" + filename

	out, err := os.Create(savePath)
	if err != nil {
		util.WriteError(w, "Gagal menyimpan gambar", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		util.WriteError(w, "Gagal menyalin gambar", http.StatusInternalServerError)
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

	if err := h.service.AddPortfolio(portfolio); err != nil {
		util.WriteError(w, "Gagal menyimpan data ke database", http.StatusInternalServerError)
		return
	}

	util.WriteSuccess(w, "Portfolio berhasil ditambahkan", nil)
}
