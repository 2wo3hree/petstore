package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(addr string, r *chi.Mux) *Server {
	return &Server{
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (g *Server) Serve() error {
	log.Println("Сервер запущен на порту", g.server.Addr)
	return g.server.ListenAndServe()
}

func (g *Server) Shutdown(ctx context.Context) error {
	log.Println("Остановка сервера...")
	return g.server.Shutdown(ctx)
}
