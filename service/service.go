package service

import "project-app-portfolio-golang-rahmadhany/repository"

type Service struct {
	User UserService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
	}
}
