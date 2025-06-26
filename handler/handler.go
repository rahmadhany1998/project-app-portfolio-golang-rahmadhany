package handler

import "project-app-portfolio-golang-rahmadhany/service"

type Handler struct {
	User *UserHandler
	// Tambahkan handler lain di sini bila perlu
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		User: NewUserHandler(s.User),
	}
}
