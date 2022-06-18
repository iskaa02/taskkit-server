package sync

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iskaa02/taskkit-server/ent"
)

type sync struct {
	client *ent.Client
}

func Routes(client *ent.Client) http.Handler {
	h := &sync{client}
	// route will be mounted on /sync
	r := chi.NewRouter()
	r.Get("/pushchanges", h.PushChanges)
	r.Get("/pullchanges", h.PullChanges)
	return r
}
