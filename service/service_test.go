package service

import (
	"project-app-portfolio-golang-rahmadhany/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	dummyRepo := &repository.Repository{}
	service := NewService(dummyRepo)
	assert.NotNil(t, service)
}
