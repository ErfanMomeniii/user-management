package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"user-management/internal/model"
	"user-management/internal/util"
)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *sqlx.DB
}

func (r UserRepository) Save(user model.User) (sql.Result, error) {
	fmt.Println(user.Id)
	q := `INSERT INTO users (id, first_name, last_name, nickname, password, email, country) VALUES (?, ?, ?, ?, ?, ?, ?)`

	uuid := util.GenerateUUId()

	return r.db.Exec(q, uuid, user.FirstName, user.LastName, user.Nickname, user.Password, user.Email, user.Country)
}

func (r UserRepository) Update(id string, user model.User) (sql.Result, error) {
	q := `UPDATE users SET first_name = ?, last_name = ?,nickname= ?,password= ?,email= ?,country= ? WHERE id = ?`

	return r.db.Exec(q, user.FirstName, user.LastName, user.Nickname, user.Password, user.Email, user.Country, id)
}

func (r UserRepository) Delete(id string) (sql.Result, error) {
	q := `DELETE FROM users WHERE id = ?`

	return r.db.Exec(q, id)
}

func (r UserRepository) Get(id string) (model.User, error) {
	var user model.User
	q := `SELECT * FROM users WHERE id = ?`

	if err := r.db.Get(&user, q, id); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r UserRepository) GetAll(page int) ([]model.User, error) {
	var users []model.User

	q := `SELECT * FROM users LIMIT ?,?`

	from, to := util.Page(page)

	if err := r.db.Select(&users, q, from, to-from); err != nil {
		if err == sql.ErrNoRows {
			return []model.User{}, nil
		}

		return nil, err
	}

	return users, nil
}

func (r UserRepository) FilterByCountry(country string, page int) ([]model.User, error) {
	var users []model.User

	q := `SELECT * FROM users WHERE country = ? LIMIT ?,?`

	from, to := util.Page(page)
	if err := r.db.Select(&users, q, country, from, to-from); err != nil {
		if err == sql.ErrNoRows {
			return []model.User{}, nil
		}

		fmt.Println(err, from, to-from)
		return nil, err
	}
	return users, nil
}
