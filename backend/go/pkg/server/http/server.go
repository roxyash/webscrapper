package httpserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	http.Server
}

func New(handler http.Handler, port string) *Server {
	return &Server{
		http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: handler,
		},
	}
}

func (s *Server) Start() {
	log.Printf("Server started at %s", s.Addr)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %s", err)
	}
}

func (s *Server) ListenQuitChanBlocking(quit <-chan os.Signal) {
	code := <-quit
	log.Printf("Recieve shotdown signal: %s", code)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
}
