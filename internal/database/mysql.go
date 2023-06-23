package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/erfanmomeniii/user-management/internal/config"
)

func InitMySQL(cfg *config.Config) (*sqlx.DB, error) {
	c := cfg.Database

	db, err := sqlx.Open(c.Driver, c.DSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(c.MaxConn)
	db.SetMaxIdleConns(c.IdleConn)
	db.SetConnMaxLifetime(c.Timeout)

	return db, nil
}
