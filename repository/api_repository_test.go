package repository_test

import (
	"database/sql"
	"testing"

	"project-app-portfolio-golang-rahmadhany/model"
	"project-app-portfolio-golang-rahmadhany/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repository.ApiRepository) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	repo := repository.NewApiRepository(db)
	return db, mock, repo
}

func TestFindFirst(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "job", "photo", "description"}).
		AddRow(1, "John Doe", "Engineer", "john.png", "Experienced engineer")

	mock.ExpectQuery(`SELECT id, name, job, photo, description FROM users LIMIT 1`).
		WillReturnRows(rows)

	user, err := repo.FindFirst()
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
}

func TestFindAllPortfolios(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "image"}).
		AddRow(1, "Project A", "a.jpg").
		AddRow(2, "Project B", "b.jpg")

	mock.ExpectQuery(`SELECT id, title, image FROM portfolios`).
		WillReturnRows(rows)

	portfolios, err := repo.FindAllPortfolios()
	assert.NoError(t, err)
	assert.Len(t, portfolios, 2)
}

func TestFindPortfolioByID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id", "title", "image", "short_description", "client", "website", "long_description"}).
		AddRow(1, "Project A", "a.jpg", "Short", "Client", "http://site.com", "Long")

	mock.ExpectQuery(`SELECT id, title, image, short_description, client, website, long_description FROM portfolios WHERE id = \$1`).
		WithArgs(1).WillReturnRows(rows)

	p, err := repo.FindPortfolioByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Project A", p.Title)
}

func TestInsertPortfolio(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	p := model.Portfolio{
		Title:            "New Project",
		Image:            "img.png",
		ShortDescription: "Short",
		Client:           "Client",
		Website:          "http://site.com",
		LongDescription:  "Long",
	}

	mock.ExpectExec("INSERT INTO portfolios").
		WithArgs(p.Title, p.Image, p.ShortDescription, p.Client, p.Website, p.LongDescription).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.InsertPortfolio(p)
	assert.NoError(t, err)
}

func TestSaveContact(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	c := model.Contact{
		Name:    "Jane",
		Email:   "jane@mail.com",
		Subject: "Hello",
		Message: "World",
	}

	mock.ExpectExec("INSERT INTO contacts").
		WithArgs(c.Name, c.Email, c.Subject, c.Message).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.SaveContact(c)
	assert.NoError(t, err)
}

func TestGetAllExperiences(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "year", "company", "position", "task"}).
		AddRow(1, 2020, "ABC Corp", "Dev", "Coding").
		AddRow(2, 2021, "XYZ Inc", "Lead", "Leading")

	mock.ExpectQuery(`SELECT id, year, company, position, task FROM experiences ORDER BY year DESC`).
		WillReturnRows(rows)

	exps, err := repo.GetAllExperiences()
	assert.NoError(t, err)
	assert.Len(t, exps, 2)
}
