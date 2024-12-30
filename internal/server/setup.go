package server

import "net/http"

func NewServer(addr string, h http.Handler) *http.Server {
	srv := http.Server{
		Addr:    addr,
		Handler: h,
	}

	return &srv
}
