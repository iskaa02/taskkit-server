package sync

import (
	"context"

	"github.com/iskaa02/taskkit-server/ent"
)

func applyTaskChanges(tc taskChanges, client *ent.Tx) error {
	ctx := context.Background()
	var err error
	for _, f := range tc.Created {
		err = f.Update(client, ctx)
		if err != nil {
			return err
		}
	}
	for _, f := range tc.Updated {
		err = f.Update(client, ctx)
		if err != nil {
			return err
		}
	}
	// for _, id := range tc.Deleted {
	// 	err = q.DeleteTask(context.Background(), id)
	// 	if err != nil {
	// 		// handle error
	// 		return err
	// 	}
	// }
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
	// for _, id := range lc.Deleted {
	// 	err = c.DeleteList(context.Background(), id)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return err
}
