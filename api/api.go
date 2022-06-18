package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iskaa02/taskkit-server/api/sync"
	"github.com/iskaa02/taskkit-server/ent"
)

type Server struct {
	client *ent.Client
	r      *chi.Mux
}

func NewServer(client *ent.Client) *Server {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Mount("/sync", sync.Routes(client))
	return &Server{client, r}
}

func (s Server) Run() error {
	err := http.ListenAndServe(":8080", s.r)
	if err != nil {
		return err
	}
	return nil
}
