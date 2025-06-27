package service

import "project-app-portfolio-golang-rahmadhany/repository"

type Service struct {
	Api ApiService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Api: NewApiService(r.Api),
	}
}
