package repository

import "github.com/jmoiron/sqlx"

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// UserRepository is the repository for the model.User.
type UserRepository struct {
	db *sqlx.DB
}
