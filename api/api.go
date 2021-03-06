package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iskaa02/taskkit-server/api/sync"
)

type Server struct {
	db *sql.DB
	r  *chi.Mux
}

func NewServer(db *sql.DB) *Server {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Mount("/sync", sync.Routes(db))
	return &Server{db, r}
}

func (s Server) Run() error {
	err := http.ListenAndServe(":8080", s.r)
	if err != nil {
		return err
	}
	return nil
}
