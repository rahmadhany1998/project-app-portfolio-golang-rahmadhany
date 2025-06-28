package service

import (
	"errors"
	"project-app-portfolio-golang-rahmadhany/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockApiRepo struct{}

func (m *mockApiRepo) FindFirst() (*model.User, error) {
	return &model.User{Name: "Jane Doe", Job: "Developer"}, nil
}

func (m *mockApiRepo) FindAllPortfolios() ([]model.Portfolio, error) {
	return []model.Portfolio{{ID: 1, Title: "Portfolio A", Image: "a.jpg"}}, nil
}

func (m *mockApiRepo) FindPortfolioByID(id int) (*model.Portfolio, error) {
	if id == 1 {
		return &model.Portfolio{ID: 1, Title: "Portfolio A"}, nil
	}
	return nil, errors.New("not found")
}

func (m *mockApiRepo) InsertPortfolio(p model.Portfolio) error {
	if p.Title == "" {
		return errors.New("invalid data")
	}
	return nil
}

func (m *mockApiRepo) SaveContact(c model.Contact) error {
	if c.Name == "" {
		return errors.New("invalid contact")
	}
	return nil
}

func (m *mockApiRepo) GetAllExperiences() ([]model.Experience, error) {
	return []model.Experience{{ID: 1, Company: "Company A"}}, nil
}

func TestApiService_GetProfile(t *testing.T) {
	svc := NewApiService(&mockApiRepo{})
	user, err := svc.GetProfile()
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", user.Name)
}

func TestApiService_GetPortfolios(t *testing.T) {
	svc := NewApiService(&mockApiRepo{})
	list, err := svc.GetPortfolios()
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestApiService_GetPortfolioByID(t *testing.T) {
	svc := NewApiService(&mockApiRepo{})
	p, err := svc.GetPortfolioByID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, p.ID)

	_, err = svc.GetPortfolioByID(999)
	assert.Error(t, err)
}

func TestApiService_AddPortfolio(t *testing.T) {
	svc := NewApiService(&mockApiRepo{})
	err := svc.AddPortfolio(model.Portfolio{Title: "Test"})
	assert.NoError(t, err)

	err = svc.AddPortfolio(model.Portfolio{})
	assert.Error(t, err)
}

func TestApiService_SubmitContact(t *testing.T) {
	svc := NewApiService(&mockApiRepo{})
	err := svc.SubmitContact(model.Contact{Name: "John"})
	assert.NoError(t, err)

	err = svc.SubmitContact(model.Contact{})
	assert.Error(t, err)
}

func TestApiService_GetAllExperiences(t *testing.T) {
	svc := NewApiService(&mockApiRepo{})
	exp, err := svc.GetAllExperiences()
	assert.NoError(t, err)
	assert.Len(t, exp, 1)
}
