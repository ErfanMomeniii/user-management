package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"user-management/internal/model"
	"user-management/internal/util"
)

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// UserRepository is the repository for the model.User.
type UserRepository struct {
	db *sqlx.DB
}

// Save creates a model.User into the database.
func (r UserRepository) Save(user model.User) (sql.Result, error) {
	q := `INSERT INTO users (id, first_name, last_name, nickname, password, email, country) VALUES (?, ?, ?, ?, ?, ?, ?)`

	uuid := util.GenerateUUId()

	return r.db.Exec(q, uuid, user.FirstName, user.LastName, user.Nickname, user.Password, user.Email, user.Country)
}
