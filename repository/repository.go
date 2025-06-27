package repository

import "database/sql"

type Repository struct {
	Api ApiRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Api: NewApiRepository(db),
	}
}
