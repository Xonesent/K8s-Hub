package grpcConn

import "google.golang.org/grpc"

type ConfigGrpc struct {
	GrpcClientHosts map[string]string           `envconfig:"GRPC_HOSTS" validate:"required"`
	GrpcClientPorts map[string]string           `envconfig:"GRPC_PORTS" validate:"required"`
	GrpcClientConn  map[string]*grpc.ClientConn `validate:"required"`
}
