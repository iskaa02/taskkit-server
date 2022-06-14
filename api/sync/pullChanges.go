package sync

import (
	"net/http"
)

type changes struct {
	List listChanges `json:"lists"`
	Task taskChanges `json:"tasks"`
}

func (s sync) PullChanges(w http.ResponseWriter, r *http.Request) {
}
