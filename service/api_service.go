package service

import (
	"project-app-portfolio-golang-rahmadhany/model"
	"project-app-portfolio-golang-rahmadhany/repository"
)

type ApiService interface {
	GetProfile() (*model.User, error)
	GetPortfolios() ([]model.Portfolio, error)
	GetPortfolioByID(id int) (*model.Portfolio, error)
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

func (s *apiService) GetPortfolios() ([]model.Portfolio, error) {
	return s.apiRepo.FindAllPortfolios()
}

func (s apiService) GetPortfolioByID(id int) (*model.Portfolio, error) {
	return s.apiRepo.FindPortfolioByID(id)
}
