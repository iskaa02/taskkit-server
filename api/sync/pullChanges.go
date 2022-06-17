package sync

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iskaa02/taskkit-server/db/models"
	"gopkg.in/guregu/null.v4"
)

type changes struct {
	List listChanges `json:"lists"`
	Task taskChanges `json:"tasks"`
}
type pullChangesReq struct {
	LastPulledAt null.Int `json:"lastPulledAt"`
}

func (s sync) PullChanges(w http.ResponseWriter, r *http.Request) {
	args := pullChangesReq{}
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "bad JSON format", http.StatusBadRequest)
		return
	}
	q := models.New(s.db)
	unixTS := int64(0)
	if args.LastPulledAt.Valid {
		unixTS = args.LastPulledAt.Int64
	}
	lastPulledAt := time.Unix(unixTS, 0)
	fmt.Println(lastPulledAt.Year())
	changes := getChanges(lastPulledAt, q)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(changes)
}
