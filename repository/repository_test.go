package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {
	// Create a mock database using sqlmock
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Initialize the repository
	repo := NewRepository(db)

	// Ensure the repository is not nil
	if repo == nil {
		t.Fatal("Expected repository to be non-nil")
	}

	// Ensure the ApiRepository field is properly initialized
	if repo.Api == nil {
		t.Fatal("Expected ApiRepository to be non-nil")
	}

	// Verify that the type of repo.Api is *apiRepo
	if _, ok := repo.Api.(*apiRepo); !ok {
		t.Errorf("Expected repo.Api to be of type *apiRepo, got %T", repo.Api)
	}
}
