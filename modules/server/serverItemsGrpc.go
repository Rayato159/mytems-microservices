package server

import (
	"log"
	"net"
)

func (s *grpcServer) RunItemsGrpc() {

	// modules := NewModule(s, nil, grpcServer)
	// pb.RegisterAuthServicesServer(
	// 	grpcServer,
	// 	NewGrpcServer(nil, nil),
	// )

	log.Printf("%s is starting on %s", s.cfg.App().Name(), s.cfg.App().Url())
	lis, err := net.Listen("tcp", s.cfg.App().Url())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.server.Serve(lis)
}
