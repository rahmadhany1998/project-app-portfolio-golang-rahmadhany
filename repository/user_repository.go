package repository

import (
	"database/sql"
	"project-app-portfolio-golang-rahmadhany/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	Create(user model.User) (int, error)
	Update(user model.User) error
	Delete(id int) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepo) GetByID(id int) (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	var u model.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) Create(user model.User) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO users(name, email) VALUES($1, $2) RETURNING id",
		user.Name, user.Email).Scan(&id)
	return id, err
}

func (r *userRepo) Update(user model.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name=$1, email=$2 WHERE id=$3",
		user.Name, user.Email, user.ID)
	return err
}

func (r *userRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
