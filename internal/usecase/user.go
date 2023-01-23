package usecase

import (
	"database/sql"
	"errors"
	"user-management/internal/model"
	"user-management/internal/repository"
)

var ErrorUserNotFound = errors.New("user not found")

func GetUsers(page int) ([]model.User, error) {
	users, err := repository.User.GetAll(page)

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func GetUserById(id string) (model.User, error) {
	user, err := repository.User.Get(id)

	if err == sql.ErrNoRows {
		return model.User{}, ErrorUserNotFound
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func FilterUsersByCountry(country string, page int) ([]model.User, error) {
	users, err := repository.User.FilterByCountry(country, page)

	if err != nil {
		return []model.User{}, nil
	}

	return users, nil
}

func SaveUser(user model.User) error {
	_, err := repository.User.Save(user)

	return err
}

func DeleteUser(id string) error {
	_, err := repository.User.Delete(id)

	return err
}

func UpdateUser(id string, user model.User) error {
	_, err := repository.User.Update(id, user)

	return err
}
