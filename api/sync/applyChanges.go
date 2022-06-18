package sync

import (
	"context"

	"github.com/iskaa02/taskkit-server/ent"
)

func applyTaskChanges(tc taskChanges, c *ent.Tx) error {
	ctx := context.Background()
	var err error
	for _, f := range tc.Created {
		err = f.Update(c, ctx)
		if err != nil {
			return err
		}
	}
	for _, f := range tc.Updated {
		err = f.Update(c, ctx)
		if err != nil {
			return err
		}
	}
	for _, id := range tc.Deleted {
		err = c.Task.UpdateOneID(id).SetIsDeleted(true).Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func applyListChanges(lc listChanges, c *ent.Tx) error {
	var err error
	ctx := context.Background()
	for _, f := range lc.Created {
		err = f.Update(c, ctx)
		if err != nil {
			return err
		}
	}
	for _, f := range lc.Updated {
		err = f.Update(c, ctx)
		if err != nil {
			return err
		}
	}
	for _, id := range lc.Deleted {
		err = c.List.UpdateOneID(id).SetIsDeleted(true).Exec(ctx)
		if err != nil {
			return err
		}
	}
	return err
}
