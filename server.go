package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {

	return http.ListenAndServeTLS("192.168.0.110:"+port, "cmd/certificate/bottle.crt", "cmd/certificate/bottle.key", handler)

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
