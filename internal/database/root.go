package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/erfanmomeniii/user-management/internal/config"
)

func Init(f func(*config.Config) (*sqlx.DB, error), cfg *config.Config) (*sqlx.DB, error) {
	db, err := f(cfg)

	return db, err
}

func Close(db *sqlx.DB) error {
	err := db.Close()

	if err != nil {
		return err
	}

	return nil
}
