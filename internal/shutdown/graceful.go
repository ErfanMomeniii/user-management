package shutdown

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"user-management/internal/log"
)

type closer struct {
	function func() error
	name     string
}

var (
	Context context.Context
	cancel  context.CancelFunc
	Closers []closer
)

func Init() {
	Context, cancel = context.WithCancel(context.Background())

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		s := <-signalChannel
		log.L.Info("system call", zap.String("signal", s.String()))
		cancel()
	}()
}

func AddCloser(name string, function func() error) {
	Closers = append(Closers, closer{function: function, name: name})
}

func Wait() {
	<-Context.Done()

	for _, c := range Closers {
		if err := c.function(); err != nil {
			log.L.Error("cannot close", zap.String("name", c.name), zap.Error(err))
		} else {
			log.L.Debug("closed successfully", zap.String("name", c.name))
		}
	}
}
