package server

import (
	"github.com/Rayato159/mytems-microservices/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type IServerGrpc interface {
	RunUsersGrpc()
	RunItemsGrpc()
	GrpcServer() *grpcServer
}

type grpcServer struct {
	server *grpc.Server
	cfg    config.IConfig
	db     *mongo.Client
}

func NewGrpcServer(cfg config.IConfig, db *mongo.Client) IServerGrpc {
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	return &grpcServer{
		server: server,
		cfg:    cfg,
		db:     db,
	}
}

func (s *grpcServer) GrpcServer() *grpcServer { return s }
