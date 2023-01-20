package repository

import "user-management/internal/database"

var User *UserRepository

func Init() {
	User = NewUserRepository(database.DB)
}
