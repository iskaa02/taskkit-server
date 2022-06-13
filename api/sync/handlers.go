package sync

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type sync struct {
	db *sql.DB
}

func Routes(db *sql.DB) http.Handler {
	h := &sync{db}
	// route will be mounted on /sync
	r := chi.NewRouter()
	r.Get("/pushchanges", h.PushChanges)
	r.Get("/pullchanges", h.PullChanges)
	return r
}
