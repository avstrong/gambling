package web

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/avstrong/gambling/internal/uservice"
	"github.com/gorilla/mux"
)

type Server struct {
	srv      *http.Server
	router   *mux.Router
	uService *uservice.Service
}

type Conf struct {
	ServerLogger      *log.Logger
	Host              string
	Port              string
	ReadHeaderTimeout time.Duration
}

func New(ctx context.Context, conf *Conf, uService *uservice.Service) (*Server, error) {
	r := mux.NewRouter()

	//nolint:exhaustruct
	srv := &http.Server{
		Addr:              net.JoinHostPort(conf.Host, conf.Port),
		ReadHeaderTimeout: conf.ReadHeaderTimeout,
		ErrorLog:          conf.ServerLogger,
		Handler:           r,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	server := &Server{
		srv:      srv,
		router:   r,
		uService: uService,
	}

	server.addRoutes(r)

	return server, nil
}

func (s *Server) Srv() *http.Server {
	return s.srv
}

func (s *Server) Router() *mux.Router {
	return s.router
}
