package sync

import (
	"context"

	"github.com/iskaa02/taskkit-server/db/models"
)

func applyTaskChanges(tc taskChanges, q *models.Queries) error {
	for _, f := range tc.Created {
		err := f.Create(q)
		if err != nil {
			// handle error
			return err
		}
	}
	for _, f := range tc.Updated {
		err := f.Update(q)
		if err != nil {
			// handle error
			return err
		}
	}
	for _, id := range tc.Deleted {
		err := q.DeleteTask(context.Background(), id)
		if err != nil {
			// handle error
			return err
		}
	}
	return nil
}

func applyListChanges(lc listChanges, q *models.Queries) error {
	for _, f := range lc.Created {
		err := f.Create(q)
		if err != nil {
			return err
		}
	}
	for _, f := range lc.Updated {
		err := f.Update(q)
		if err != nil {
			return err
		}
	}
	for _, id := range lc.Deleted {
		err := q.DeleteList(context.Background(), id)
		if err != nil {
			return err
		}
	}
	return nil
}
