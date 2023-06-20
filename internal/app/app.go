package app

import (
	"context"
	"github.com/erfanmomeniii/user-management/internal/config"
	"github.com/erfanmomeniii/user-management/internal/database"
	httpServer "github.com/erfanmomeniii/user-management/internal/http/server"
	"github.com/erfanmomeniii/user-management/internal/log"
	"github.com/erfanmomeniii/user-management/internal/repository"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Config   *config.Config
	Logger   *zap.Logger
	Database *sqlx.DB
}

func New(configPath string) (*App, error) {
	cfg, err := config.Init(configPath)
	if err != nil {
		return nil, err
	}

	logger, err := log.Init(cfg.Logger.Level)
	if err != nil {
		return nil, err
	}

	db, err := database.Init(database.InitMySQL, cfg)
	if err != nil {
		return nil, err
	}

	app := &App{
		Config:   cfg,
		Logger:   logger,
		Database: db,
	}

	return app, err
}

func (app *App) Start() {
	httpServer.Init(app.Config)

	repository.Init(app.Database)

	httpServer.Serve(app.Logger, app.Config)
}

func (app *App) Wait() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		s := <-signalChannel
		app.Logger.Info("system call", zap.String("signal", s.String()))
		cancel()
	}()
	return ctx
}

func (app *App) Shutdown(ctx context.Context) error {
	<-ctx.Done()

	err := httpServer.Close()
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "http server"), zap.Error(err))

		return err
	}

	err = database.Close(app.Database)
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "database"), zap.Error(err))

		return err
	}

	err = log.Close(app.Logger)
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "logger"), zap.Error(err))

		return err
	}

	return nil
}
