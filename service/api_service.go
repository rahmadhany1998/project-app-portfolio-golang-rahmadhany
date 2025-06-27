package service

import (
	"project-app-portfolio-golang-rahmadhany/model"
	"project-app-portfolio-golang-rahmadhany/repository"
)

type ApiService interface {
	GetProfile() (*model.User, error)
}

type apiService struct {
	apiRepo repository.ApiRepository
}

func NewApiService(repo repository.ApiRepository) ApiService {
	return &apiService{repo}
}

func (s apiService) GetProfile() (*model.User, error) {
	return s.apiRepo.FindFirst()
}
