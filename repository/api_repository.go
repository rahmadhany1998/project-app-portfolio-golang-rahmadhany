package repository

import (
	"database/sql"
	"project-app-portfolio-golang-rahmadhany/model"
)

type ApiRepository interface {
	FindFirst() (*model.User, error)
}

type apiRepo struct {
	db *sql.DB
}

func NewApiRepository(db *sql.DB) ApiRepository {
	return &apiRepo{db}
}

func (r apiRepo) FindFirst() (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name, job FROM users LIMIT 1")

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Job)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
