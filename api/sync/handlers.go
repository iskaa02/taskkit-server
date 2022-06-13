package sync

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type sync struct {
	db sql.DB
}

func Sync(db sql.DB) http.Handler {
	h := &sync{db}
	r := chi.NewRouter()
	r.Get("/sync/pushchanges", h.PushChanges)
	r.Get("/sync/pullchanges", h.PullChanges)
	return r
}
