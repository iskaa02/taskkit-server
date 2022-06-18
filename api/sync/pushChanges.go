package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s sync) PushChanges(w http.ResponseWriter, r *http.Request) {
	changes := changes{}
	err := json.NewDecoder(r.Body).Decode(&changes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	tx, err := s.client.Tx(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = applyListChanges(changes.List, tx)
	if err != nil {
		tx.Rollback()
		if err.Error() == "conflict" {
			http.Error(w, "Tried to modify deleted record, Pull latest changes first", http.StatusBadRequest)
			return
		}
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	err = applyTaskChanges(changes.Task, tx)
	if err != nil {
		if err.Error() == "conflict" {
			http.Error(w, "Tried to modify deleted record, Pull latest changes first", http.StatusBadRequest)
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
