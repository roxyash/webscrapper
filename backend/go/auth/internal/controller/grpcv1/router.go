package grpcv1

import (
	"github.com/roxyash/webscrapper_proto/gen/go/auth"
	"google.golang.org/grpc"
	"webscrapper/auth/internal/service"
	grpcserver "webscrapper/pkg/server/grpc"
)

type Router struct {
	Auth auth.AuthServiceServer
}

func (r *Router) RegistrationFunc() grpcserver.RegisterToServer {
	return func(srv *grpc.Server) {
		auth.RegisterAuthServiceServer(srv, r.Auth)
	}
}

func New(services *service.Service) *Router {
	return &Router{
		Auth: NewAuth(services.Auth),
	}
}
