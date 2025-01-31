package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/steliosmagalios/link-shortener/internal/server/middleware"
)

type Server struct {
	srv *http.Server
}

func New(addr string, h http.Handler) *Server {
	// Setup Middleware
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.EnableCORS,
	)

	srv := http.Server{
		Addr:    addr,
		Handler: stack(h),
	}

	return &Server{
		srv: &srv,
	}
}

func (s *Server) Start() {
	log.Printf("Server starting and listening on port %v", s.srv.Addr)
	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}
}
