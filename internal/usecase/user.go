package usecase

import "user-management/internal/model"

func GetUsers(page int) []model.User {
	
	return []model.User{}
}

func GetUserById(id int) (model.User, error) {
	return model.User{}, nil
}

func FilterUsersByCountry(country string, page int) []model.User {
	return []model.User{}
}

func SaveUser(user model.User) error {
	return nil
}

func DeleteUser(id int) error {
	return nil
}

func UpdateUser(user model.User) error {
	return nil
}
