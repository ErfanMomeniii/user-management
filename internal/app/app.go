package app

import (
	"user-management/internal/config"
	"user-management/internal/database"
	httpServer "user-management/internal/http/server"
	"user-management/internal/log"
	"user-management/internal/repository"
	"user-management/internal/shutdown"
)

func Init(configPath string) (err error) {
	if err = config.Init(configPath); err != nil {
		return err
	}

	if err = log.Init(config.C.Logger.Level); err != nil {
		return err
	}

	if err = database.Init(database.InitMySQL); err != nil {
		return err
	}

	httpServer.Init()

	repository.Init()

	shutdown.Init()

	shutdown.AddCloser("log", log.Close)
	shutdown.AddCloser("httpServer", httpServer.Close)
	shutdown.AddCloser("database", database.Close)

	return err
}
