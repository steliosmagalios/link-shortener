package server

import "net/http"

func NewServer(addr string) *http.Server {
	srv := http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(http.NotFound),
	}

	return &srv
}
