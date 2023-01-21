package database

import (
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(f func() (*sqlx.DB, error)) (err error) {
	DB, err = f()
	return err
}

func Close() error {
	err := DB.Close()

	if err != nil {
		return err
	}

	return nil
}
