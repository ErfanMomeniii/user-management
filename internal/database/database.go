package database

import (
	"github.com/jmoiron/sqlx"
)

// DB is the global instance of database.
var DB *sqlx.DB

// Init creates an instance of the database based on the config.
func Init(f func() (*sqlx.DB, error)) (err error) {
	DB, err = f()
	return err
}
