package sync

import (
	"context"
	"fmt"
	"time"

	"github.com/iskaa02/taskkit-server/ent"
	"github.com/iskaa02/taskkit-server/ent/list"
	"github.com/iskaa02/taskkit-server/ent/task"
)

func getChanges(lastPulled time.Time, c *ent.Client) changes {
	return changes{
		List: getListChanges(lastPulled, c),
		Task: getTaskChanges(lastPulled, c),
	}
}

func getTaskChanges(lastPulled time.Time, c *ent.Client) taskChanges {
	ctx := context.Background()
	rawCreated, err := c.Task.Query().Where(task.CreatedAtGTE(lastPulled)).All(ctx)
	if err != nil {
		fmt.Println(err)
	}
	created := make([]rawTask, len(rawCreated))
	for i := range rawCreated {
		created[i] = toRawTask(rawCreated[i])
	}

	rawUpdated, err := c.Task.Query().
		Where(
			task.And(
				task.CreatedAtGT(lastPulled),
				task.LastModifiedGTE(lastPulled),
			),
		).
		All(ctx)
	if err != nil {
		fmt.Println(err)
	}
	updated := make([]rawTask, len(rawUpdated))
	for i := range rawUpdated {
		updated[i] = toRawTask(rawUpdated[i])
	}

	deleted, err := c.Task.Query().
		Where(task.LastModifiedGTE(lastPulled)).
		IDs(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return taskChanges{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}
}

func getListChanges(lastPulled time.Time, c *ent.Client) listChanges {
	ctx := context.Background()
	rawCreated, err := c.List.Query().Where(list.CreatedAtGTE(lastPulled)).All(ctx)
	if err != nil {
		fmt.Println(err)
	}
	created := make([]rawList, len(rawCreated))
	for i := range rawCreated {
		created[i] = toRawList(rawCreated[i])
	}
	rawUpdated, err := c.List.Query().
		Where(
			list.And(
				list.CreatedAtGT(lastPulled),
				list.LastModifiedGTE(lastPulled),
			),
		).
		All(ctx)
	if err != nil {
		fmt.Println(err)
	}
	updated := make([]rawList, len(rawUpdated))
	for i := range rawUpdated {
		updated[i] = toRawList(rawUpdated[i])
	}

	deleted, err := c.List.Query().
		Where(list.LastModifiedGTE(lastPulled)).
		IDs(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return listChanges{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}
}

func toRawTask(t *ent.Task) rawTask {
	return rawTask{
		ID:          t.ID,
		Description: t.Description,
		Name:        t.Name,
		Subtasks:    *t.Subtasks,
		Repeat:      t.Repeat,
		Reminder:    t.Reminder,
		ListID:      t.ListID,
		IsCompleted: t.IsCompleted,
	}
}

func toRawList(l *ent.List) rawList {
	theme, err := l.QueryTheme().Only(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return rawList{
		Name: l.Name,
		ID:   l.ID,
		Theme: Theme{
			Primary:   theme.Primary,
			Secondary: theme.Secondary,
		},
	}
}
