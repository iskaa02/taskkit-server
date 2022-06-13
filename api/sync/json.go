package sync

import (
	"context"

	"github.com/iskaa02/taskkit-server/db/models"
)

type (
	rawTask models.Task

	rawList struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Theme struct {
			Primary   string `json:"primary"`
			Secondary string `json:"secondary"`
		} `json:"theme"`
	}
	listChanges struct {
		Created []rawList `json:"created"`
		Updated []rawList `json:"updated"`
		Deleted []string  `json:"deleted"`
	}
	taskChanges struct {
		Created []rawTask `json:"created"`
		Updated []rawTask `json:"updated"`
		Deleted []string  `json:"deleted"`
	}
)

func (t rawTask) Create(q *models.Queries) error {
	err := q.CreateTask(context.Background(), models.CreateTaskParams{
		ID:          t.ID,
		Name:        t.Name,
		Subtasks:    t.Subtasks,
		ListID:      t.ListID,
		Description: t.Description,
		Reminder:    t.Reminder,
		Repeat:      t.Repeat,
		IsCompleted: t.IsCompleted,
	})
	return err
}

func (t rawTask) Update(q *models.Queries) error {
	err := q.UpdateTask(context.Background(), models.UpdateTaskParams{
		ID:          t.ID,
		Name:        t.Name,
		Subtasks:    t.Subtasks,
		ListID:      t.ListID,
		Description: t.Description,
		Reminder:    t.Reminder,
		Repeat:      t.Repeat,
		IsCompleted: t.IsCompleted,
	})
	return err
}

func (l rawList) Create(q *models.Queries) error {
	themeID, err := q.FindTheme(context.Background(), models.FindThemeParams{
		Primary: l.Theme.Primary, Secondary: l.Theme.Secondary,
	})
	if err != nil {
		return err
	}
	err = q.CreateList(context.Background(), models.CreateListParams{
		ID:      l.ID,
		Name:    l.Name,
		ThemeID: themeID,
	})
	return err
}

func (l rawList) Update(q *models.Queries) error {
	themeID, err := q.FindTheme(context.Background(), models.FindThemeParams{
		Primary: l.Theme.Primary, Secondary: l.Theme.Secondary,
	})
	if err != nil {
		return err
	}
	err = q.UpdateList(context.Background(), models.UpdateListParams{
		ID:      l.ID,
		Name:    l.Name,
		ThemeID: themeID,
	})
	return err
}
