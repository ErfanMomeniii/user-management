package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
	"user-management/internal/config"
	internalHandler "user-management/internal/http/handler"
	"user-management/internal/http/validator"
	"user-management/internal/log"
)

var E *echo.Echo

func Init() {
	E = echo.New()

	E.HideBanner = true
	E.Server.ReadTimeout = config.C.HTTPServer.ReadTimeout
	E.Server.WriteTimeout = config.C.HTTPServer.WriteTimeout
	E.Server.ReadHeaderTimeout = config.C.HTTPServer.ReadHeaderTimeout
	E.Server.IdleTimeout = config.C.HTTPServer.IdleTimeout
	E.Validator = validator.New()
}

func Serve() {
	v1 := E.Group("/v1")
	{
		u := v1.Group("/user")
		{
			u.POST("/", internalHandler.SaveUser)
			u.GET("/", internalHandler.GetAllUser)
			u.DELETE("/:userId", internalHandler.DeleteUser)
			u.PUT("/:userId", internalHandler.UpdateUser)
		}
	}

	go func() {
		if err := E.Start(config.C.HTTPServer.Listen); err != nil && err != http.ErrServerClosed {
			log.L.Fatal(
				"cannot start the server", zap.String("listen", config.C.HTTPServer.Listen), zap.Error(err),
			)
		}
	}()
}

func Close() error {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return E.Shutdown(c)
}
