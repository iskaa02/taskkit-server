package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/iskaa02/taskkit-server/ent"
)

func (s sync) PushChanges(w http.ResponseWriter, r *http.Request) {
	changes := changes{}
	err := json.NewDecoder(r.Body).Decode(&changes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err != nil {
		panic(err)
	}
	// queries := models.New(tx)
	entDriver := entsql.OpenDB("postgres", s.db)
	client := ent.NewClient(ent.Driver(entDriver))
	tx, _ := client.Tx(context.Background())

	err = applyListChanges(changes.List, tx)
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
	err = applyTaskChanges(changes.Task, tx)
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
