package repository

import (
	"database/sql"

	"github.com/erfanmomeniii/user-management/internal/model"
)

type UserRepositoryMock struct {
	ExpErrSave               error
	ExpErrUpdate             error
	ExpErrDelete             error
	ExpErrGet                error
	ExpErrGetAll             error
	ExpErrFindByCountry      error
	ExpResulGet              model.User
	ExpResultGetAll          []model.User
	ExpResultFilterByCountry []model.User
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (u UserRepositoryMock) Save(_ model.User) (sql.Result, error) {
	return nil, u.ExpErrSave
}

func (u UserRepositoryMock) Update(_ string, _ model.User) (sql.Result, error) {
	return nil, u.ExpErrUpdate
}

func (u UserRepositoryMock) Delete(_ string) (sql.Result, error) {
	return nil, u.ExpErrDelete
}

func (u UserRepositoryMock) Get(_ string) (model.User, error) {
	return u.ExpResulGet, u.ExpErrGet
}

func (u UserRepositoryMock) GetAll(_ int) ([]model.User, error) {
	return u.ExpResultGetAll, u.ExpErrGetAll
}

func (u UserRepositoryMock) FilterByCountry(_ string, _ int) ([]model.User, error) {
	return u.ExpResultFilterByCountry, u.ExpErrFindByCountry
}
