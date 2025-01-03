package server

import (
	"errors"
	"log"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func New(addr string, h http.Handler) *Server {
	srv := http.Server{
		Addr:    addr,
		Handler: h,
	}

	return &Server{
		srv: &srv,
	}
}

func (s *Server) Start() {
	log.Printf("Server starting and listening on http://%v", s.srv.Addr)
	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}

}
