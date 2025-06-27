package handler

import (
	"net/http"
	"project-app-portfolio-golang-rahmadhany/service"
	"project-app-portfolio-golang-rahmadhany/util"
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
