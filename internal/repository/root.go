package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/erfanmomeniii/user-management/internal/model"
)

var User UserDatabaseOperation

func Init(db *sqlx.DB) {
	User = NewUserRepository(db)
}

type UserDatabaseOperation interface {
	Save(model.User) (sql.Result, error)
	Update(string, model.User) (sql.Result, error)
	Delete(string) (sql.Result, error)
	Get(string) (model.User, error)
	GetAll(int) ([]model.User, error)
	FilterByCountry(string, int) ([]model.User, error)
}
