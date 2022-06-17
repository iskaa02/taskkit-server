package sync

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iskaa02/taskkit-server/db/models"
)

func (s sync) PushChanges(w http.ResponseWriter, r *http.Request) {
	changes := changes{}
	err := json.NewDecoder(r.Body).Decode(&changes)
	if err != nil {
		http.Error(w, "bad JSON format", http.StatusBadRequest)
	}
	tx, err := s.db.Begin()
	if err != nil {
		panic(err)
	}
	queries := models.New(tx)
	err = applyListChanges(changes.List, queries)
	if err != nil {
		tx.Rollback()
		if err.Error() == "conflict" {
			http.Error(w, "Tried to update deleted record, Pull latest changes first", http.StatusBadRequest)
			return
		}
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	err = applyTaskChanges(changes.Task, queries)
	if err != nil {
		if err.Error() == "conflict" {
			http.Error(w, "Tried to update deleted record, Pull latest changes first", http.StatusBadRequest)
			return
		}
		tx.Rollback()
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	tx.Commit()
	w.WriteHeader(200)
}
