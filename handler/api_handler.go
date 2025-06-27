package handler

import (
	"net/http"
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
