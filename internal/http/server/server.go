package server

import (
	"context"
	"net/http"
	"time"

	"github.com/erfanmomeniii/user-management/internal/config"
	internalHandler "github.com/erfanmomeniii/user-management/internal/http/handler"
	"github.com/erfanmomeniii/user-management/internal/http/validator"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var E *echo.Echo

func Init(cfg *config.Config) {
	E = echo.New()

	E.HideBanner = true
	E.Server.ReadTimeout = cfg.HTTPServer.ReadTimeout
	E.Server.WriteTimeout = cfg.HTTPServer.WriteTimeout
	E.Server.ReadHeaderTimeout = cfg.HTTPServer.ReadHeaderTimeout
	E.Server.IdleTimeout = cfg.HTTPServer.IdleTimeout
	E.Validator = validator.New()
}

func Serve(log *zap.Logger, cfg *config.Config) {
	v1 := E.Group("/v1")
	{
		v1.POST("/user", internalHandler.SaveUser)
		v1.GET("/users", internalHandler.GetUsers)
		v1.GET("/user/:userId", internalHandler.GetUser)
		v1.DELETE("/user/:userId", internalHandler.DeleteUser)
		v1.PUT("/user/:userId", internalHandler.UpdateUser)
	}

	go func() {
		if err := E.Start(cfg.HTTPServer.Listen); err != nil && err != http.ErrServerClosed {
			log.Fatal(
				"cannot start the server", zap.String("listen", cfg.HTTPServer.Listen), zap.Error(err),
			)
		}
	}()
}

func Close() error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return E.Shutdown(c)
}
