package repository

import (
	"database/sql"
	"user-management/internal/database"
	"user-management/internal/model"
)

var User UserDatabaseOperation

func Init() {
	User = NewUserRepository(database.DB)
}

type UserDatabaseOperation interface {
	Save(model.User) (sql.Result, error)
	Update(string, model.User) (sql.Result, error)
	Delete(string) (sql.Result, error)
	Get(string) (model.User, error)
	GetAll(int) ([]model.User, error)
	FilterByCountry(string, int) ([]model.User, error)
}
