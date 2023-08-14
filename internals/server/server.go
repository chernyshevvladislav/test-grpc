package server

import (
	"github.com/chernyshevvladislav/test-grpc/internals/config"
	"github.com/chernyshevvladislav/test-grpc/internals/delivery/grpc"
	"github.com/chernyshevvladislav/test-grpc/pkg/generate/library"
	grpcServer "google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	Server   *grpcServer.Server
	Listener net.Listener
}

func NewServer(cfg config.HTTP, handler *grpc.Handler) *Server {
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalf("Cannot start listener: %s", err)
	}
	server := grpcServer.NewServer()
	library.RegisterLibraryServer(server, handler)
	return &Server{Server: server, Listener: lis}
}

func (s *Server) Start() error {
	err := s.Server.Serve(s.Listener)
	if err != nil {
		return err
	}

	return nil
}
