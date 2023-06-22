package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/erfanmomeniii/user-management/internal/config"
	"github.com/erfanmomeniii/user-management/internal/database"
	grpcServer "github.com/erfanmomeniii/user-management/internal/grpc/server"
	httpServer "github.com/erfanmomeniii/user-management/internal/http/server"
	"github.com/erfanmomeniii/user-management/internal/log"
	"github.com/erfanmomeniii/user-management/internal/repository"
	"github.com/erfanmomeniii/user-management/internal/tracing"

	"github.com/jmoiron/sqlx"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var (
	traceProvider *traceSdk.TracerProvider
)

type App struct {
	Config   *config.Config
	Logger   *zap.Logger
	Database *sqlx.DB
	Tracer   trace.Tracer
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

	provider, tracer, err := tracing.Init(tracing.InitJaeger, cfg)
	if err != nil {
		return nil, err
	}

	traceProvider = provider

	app := &App{
		Config:   cfg,
		Logger:   logger,
		Database: db,
		Tracer:   tracer,
	}

	return app, err
}

func (app *App) Start() {
	httpServer.Init(app.Config)
	grpcServer.Init(app.Config)

	repository.Init(app.Database)

	httpServer.Serve(app.Logger, app.Config)
	grpcServer.Serve(app.Logger, app.Config)
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

	err = grpcServer.Close()
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "grpc server"), zap.Error(err))

		return err
	}

	err = database.Close(app.Database)
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "database"), zap.Error(err))

		return err
	}

	err = tracing.Close(traceProvider)
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "tracing"), zap.Error(err))

		return err
	}

	err = log.Close(app.Logger)
	if err != nil {
		app.Logger.Error("cannot close", zap.String("name", "logger"), zap.Error(err))

		return err
	}

	return nil
}
