package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type server struct {
	mux        *http.ServeMux
	httpServer *http.Server
}

type Option func(*server)

// WithReadTimeout sets the read timeout on the
// underlying HttpServer. If unset, the default is
// 30s.
func WithReadTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.httpServer.ReadTimeout = timeout
	}
}

// WithWriteTimeout sets the write timeout on the
// underlying HttpServer. If unset, the default
// is 30s.
func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.httpServer.WriteTimeout = timeout
	}
}

// WithAddress sets the address on the underlying
// HttpServer. If unset, the default is ":8080".
func WithAddress(addr string) Option {
	return func(s *server) {
		s.httpServer.Addr = addr
	}
}

var defaultHttpServer = &http.Server{
	WriteTimeout: time.Second * 30,
	ReadTimeout:  time.Second * 30,
	Addr:         ":8080",
}

// NewServer returns a server with the provided
// timeout options after registering the handlers.
// Default timeouts are 30s and the default address
// is ":8080".
func NewServer(opts ...Option) *server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/resources", getResourceHandler)
	s := &server{
		mux:        mux,
		httpServer: defaultHttpServer,
	}
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Run is a blocking function that starts the
// HTTP server. Must call server.Stop to gracefully
// shut down a running server.
func (s *server) Run() error {
	log.Println("Running server")
	defer log.Print("server run ended")
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop gracefully shuts down the HTTP server or returns
// an error if unable to do so.
func (s *server) Stop(ctx context.Context) error {
	log.Println("Stopping server")
	defer log.Println("server shut down")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
