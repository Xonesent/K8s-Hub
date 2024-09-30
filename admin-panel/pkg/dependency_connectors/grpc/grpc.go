package grpcServer

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type GRPCConfig struct {
	Host string `envconfig:"GRPC_HOST" validate:"required"`
	Port string `envconfig:"GRPC_PORT" validate:"required"`
}

func NewGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			loggingInterceptor(),
			errorHandlingInterceptor(),
		),
	)

	//grpc_health_v1.RegisterHealthServer(grpcServer, service.NewHealthService())

	reflection.Register(grpcServer)

	return grpcServer
}

func loggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		zap.L().Info("Received gRPC request",
			zap.String("method", info.FullMethod),
		)

		resp, err := handler(ctx, req)

		if err != nil {
			zap.L().Error("gRPC request failed", zap.String("method", info.FullMethod), zap.Error(err))
		} else {
			zap.L().Info("gRPC request successful", zap.String("method", info.FullMethod))
		}

		return resp, err
	}
}

func errorHandlingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, status.Errorf(status.Code(err), "gRPC error: %v", err)
		}
		return resp, nil
	}
}
