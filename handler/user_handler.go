package handler

import (
	"encoding/json"
	"net/http"
	"project-app-portfolio-golang-rahmadhany/dto"
	"project-app-portfolio-golang-rahmadhany/service"
	"project-app-portfolio-golang-rahmadhany/util"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		util.WriteError(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	util.WriteSuccess(w, "List of users", users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.service.GetByID(id)
	if err != nil {
		util.WriteError(w, "User not found", http.StatusNotFound)
		return
	}
	util.WriteSuccess(w, "User found", user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.service.Create(req)
	if err != nil {
		util.WriteError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	util.WriteSuccess(w, "User created", map[string]int{"id": id})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.Update(id, req); err != nil {
		util.WriteError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	util.WriteSuccess(w, "User updated", nil)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		util.WriteError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	util.WriteSuccess(w, "User deleted", nil)
}
