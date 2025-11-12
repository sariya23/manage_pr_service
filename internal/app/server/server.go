package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Srv *http.Server
}

func NewServer(serverAddress string, serverPort int, handler http.Handler) *Server {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverAddress, serverPort),
		Handler: handler,
	}
	return &Server{Srv: server}
}

func (srv *Server) MustRun() {
	if err := srv.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
