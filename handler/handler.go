// File: handler/handler.go
package handler

import "project-app-portfolio-golang-rahmadhany/service"

type Handler struct {
	Api      *ApiHandler
	Frontend *FrontendHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Api:      NewApiHandler(s.Api),
		Frontend: NewFrontendHandler(s.Api),
	}
}
