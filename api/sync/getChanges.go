package sync

import (
	"context"
	"fmt"
	"time"

	"github.com/iskaa02/taskkit-server/db/models"
	"gopkg.in/guregu/null.v4"
)

func getChanges(lastPulled time.Time, q *models.Queries) changes {
	return changes{
		List: getListChanges(lastPulled, q),
		Task: getTaskChanges(lastPulled, q),
	}
}

func getTaskChanges(lastPulled time.Time, q *models.Queries) taskChanges {
	ctx := context.Background()
	rawCreated, err := q.GetNewlyCreatedTasks(ctx, lastPulled)
	if err != nil {
		fmt.Println(err)
	}
	created := make([]rawTask, len(rawCreated))
	for i := range rawCreated {
		created[i] = rawTask(rawCreated[i])
	}

	rawUpdated, err := q.GetNewlyUpdatedTasks(ctx, lastPulled)
	if err != nil {
		fmt.Println(err)
	}
	updated := make([]rawTask, len(rawUpdated))
	for i := range rawUpdated {
		updated[i] = rawTask(rawUpdated[i])
	}

	deleted, err := q.GetNewlyDeletedTasks(ctx, lastPulled)
	if err != nil {
		fmt.Println(err)
	}
	return taskChanges{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}
}

func getListChanges(lastPulled time.Time, q *models.Queries) listChanges {
	ctx := context.Background()
	rawCreated, err := q.GetNewlyCreatedLists(ctx, lastPulled)
	if err != nil {
		fmt.Println(err)
	}
	created := make([]rawList, len(rawCreated))
	for i := range rawCreated {
		created[i] = toRawList(rawCreated[i])
	}
	rawUpdated, err := q.GetNewlyUpdatedLists(ctx, lastPulled)
	if err != nil {
		fmt.Println(err)
	}
	updated := make([]rawList, len(rawUpdated))
	for i := range rawUpdated {
		updated[i] = toRawList(models.GetNewlyCreatedListsRow(rawUpdated[i]))
	}

	deleted, err := q.GetNewlyDeletedLists(ctx, lastPulled)
	if err != nil {
		fmt.Println(err)
	}
	return listChanges{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}
}

func toRawList(l models.GetNewlyCreatedListsRow) rawList {
	return rawList{
		Name: l.Name,
		ID:   l.ID,
		Theme: Theme{
			Primary:   l.Primary,
			Secondary: null.StringFrom(l.Secondary),
		},
	}
}
