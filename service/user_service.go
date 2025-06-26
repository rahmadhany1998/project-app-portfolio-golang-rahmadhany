package service

import (
	"project-app-portfolio-golang-rahmadhany/dto"
	"project-app-portfolio-golang-rahmadhany/model"
	"project-app-portfolio-golang-rahmadhany/repository"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	Create(input dto.CreateUserRequest) (int, error)
	Update(id int, input dto.UpdateUserRequest) error
	Delete(id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.userRepo.GetAll()
}

func (s *userService) GetByID(id int) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *userService) Create(input dto.CreateUserRequest) (int, error) {
	return s.userRepo.Create(model.User{
		Name:  input.Name,
		Email: input.Email,
	})
}

func (s *userService) Update(id int, input dto.UpdateUserRequest) error {
	return s.userRepo.Update(model.User{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	})
}

func (s *userService) Delete(id int) error {
	return s.userRepo.Delete(id)
}
