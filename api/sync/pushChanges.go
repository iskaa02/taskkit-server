package sync

import (
	"encoding/json"
	"net/http"

	"github.com/iskaa02/taskkit-server/db/models"
)

func (s sync) PushChanges(w http.ResponseWriter, r *http.Request) {
	changes := changes{}
	json.NewDecoder(r.Body).Decode(&changes)
	tx, err := s.db.Begin()
	if err != nil {
		panic(err)
	}
	queries := models.New(tx)
	err = applyListChanges(changes.List, queries)
	err = applyTaskChanges(changes.Task, queries)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx.Commit()
	w.WriteHeader(200)
}
