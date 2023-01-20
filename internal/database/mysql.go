package database

import (
	"github.com/jmoiron/sqlx"
	"user-management/internal/config"
)

// InitMySQL creates a new db instance connected to the MySQL.
func InitMySQL() (*sqlx.DB, error) {
	c := config.C.Database

	db, err := sqlx.Open(c.Driver, c.DSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.MaxConn)
	db.SetMaxIdleConns(c.IdleConn)
	db.SetConnMaxLifetime(c.Timeout)

	return db, nil
}
