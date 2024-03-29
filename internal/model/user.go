package model

import "time"

type User struct {
	Id        string    `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Nickname  string    `db:"nickname"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	Country   string    `db:"country"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
