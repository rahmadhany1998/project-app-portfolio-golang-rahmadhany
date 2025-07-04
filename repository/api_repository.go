package repository

import (
	"database/sql"
	"project-app-portfolio-golang-rahmadhany/model"
)

type ApiRepository interface {
	FindFirst() (*model.User, error)
	FindAllPortfolios() ([]model.Portfolio, error)
	FindPortfolioByID(id int) (*model.Portfolio, error)
	InsertPortfolio(p model.Portfolio) error
	SaveContact(contact model.Contact) error
	GetAllExperiences() ([]model.Experience, error)
}

type apiRepo struct {
	db *sql.DB
}

func NewApiRepository(db *sql.DB) ApiRepository {
	return &apiRepo{db}
}

func (r apiRepo) FindFirst() (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name, job, photo, description FROM users LIMIT 1")

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Job, &user.Photo, &user.Description)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *apiRepo) FindAllPortfolios() ([]model.Portfolio, error) {
	rows, err := r.db.Query("SELECT id, title, image FROM portfolios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var portfolios []model.Portfolio
	for rows.Next() {
		var p model.Portfolio
		if err := rows.Scan(&p.ID, &p.Title, &p.Image); err != nil {
			return nil, err
		}
		portfolios = append(portfolios, p)
	}
	return portfolios, nil
}

func (r apiRepo) FindPortfolioByID(id int) (*model.Portfolio, error) {
	row := r.db.QueryRow(`
		SELECT id, title, image, short_description, client, website, long_description
		FROM portfolios
		WHERE id = $1
	`, id)

	var p model.Portfolio
	err := row.Scan(&p.ID, &p.Title, &p.Image, &p.ShortDescription, &p.Client, &p.Website, &p.LongDescription)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r apiRepo) InsertPortfolio(p model.Portfolio) error {
	query := `
		INSERT INTO portfolios (title, image, short_description, client, website, long_description)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, p.Title, p.Image, p.ShortDescription, p.Client, p.Website, p.LongDescription)
	return err
}

func (r apiRepo) SaveContact(contact model.Contact) error {
	_, err := r.db.Exec(`
		INSERT INTO contacts (name, email, subject, message)
		VALUES ($1, $2, $3, $4)`,
		contact.Name, contact.Email, contact.Subject, contact.Message)
	return err
}

func (r apiRepo) GetAllExperiences() ([]model.Experience, error) {
	rows, err := r.db.Query("SELECT id, year, company, position, task FROM experiences ORDER BY year DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []model.Experience
	for rows.Next() {
		var exp model.Experience
		if err := rows.Scan(&exp.ID, &exp.Year, &exp.Company, &exp.Position, &exp.Task); err != nil {
			return nil, err
		}
		experiences = append(experiences, exp)
	}
	return experiences, nil
}
