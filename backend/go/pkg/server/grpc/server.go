package grpcserver

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type RegisterToServer func(srv *grpc.Server)

type InitializedServer interface {
	RegistrationFunc() RegisterToServer
}

type Server struct {
	srv    *grpc.Server
	listen net.Listener
}

func New(port string, router InitializedServer) (*Server, error) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	srv := grpc.NewServer()

	registerToServer := router.RegistrationFunc()
	registerToServer(srv)

	return &Server{srv: srv, listen: listen}, nil
}

func (s *Server) Start() {
	log.Printf("gRPC Server started at %s", s.listen.Addr().String())
	if err := s.srv.Serve(s.listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *Server) ListenQuitChanBlocking(quit <-chan os.Signal) {
	<-quit // Received shutdown signal
	log.Println("Received an interrupt signal, stop the service...")
	if err := s.Shutdown(); err != nil {
		log.Printf("Error while stopping Server: %s", err)
	}
	os.Exit(0)
}

func (s *Server) Shutdown() error {
	s.srv.GracefulStop()
	log.Printf("Server gracefully stopped")
	return nil
}
