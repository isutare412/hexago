package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/isutare412/hexago/gateway/pkg/config"
	"github.com/isutare412/hexago/gateway/pkg/core/port"
)

type Server struct {
	srv *http.Server
}

func NewServer(
	cfg *config.HttpServerConfig,
	uSvc port.UserService,
) *Server {
	root := mux.NewRouter()
	root.Use(wrapResponseWriter, accessLog)

	api := root.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", createUser(uSvc)).Methods("POST")
	api.HandleFunc("/users", getUser(uSvc)).Methods("GET")
	api.HandleFunc("/users", deleteUser(uSvc)).Methods("DELETE")

	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: root,
		},
	}
}

func (s *Server) Run(ctx context.Context) <-chan error {
	fails := make(chan error, 1)
	go func() {
		defer close(fails)

		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fails <- fmt.Errorf("listening on http server: %w", err)
			return
		}
	}()
	return fails
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
