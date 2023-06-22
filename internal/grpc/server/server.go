package server

import (
	"github.com/erfanmomeniii/user-management/internal/config"
	"github.com/erfanmomeniii/user-management/internal/grpc/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var (
	s        *grpc.Server
	listener net.Listener
)

func Init(cfg *config.Config) {
	s = grpc.NewServer(
		grpc.ConnectionTimeout(cfg.GRPCServer.ConnectionTimeout))
}

func Serve(log *zap.Logger, cfg *config.Config) {
	listener, _ = net.Listen("tcp", cfg.GRPCServer.Address)

	proto.RegisterUserServer(s, proto.UserServer{})

	go func() {
		log.Info("Starting gRPC server", zap.String("address", cfg.GRPCServer.Address))
		if err := s.Serve(listener); err != nil {
			log.Fatal("Could not serve gRPC server: %v", zap.Error(err))
		}
	}()
}

func Close() error {
	err := listener.Close()
	if err != nil {
		return err
	}
	s.GracefulStop()

	return nil
}
