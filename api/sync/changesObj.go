package sync

import (
	"context"
	"database/sql"

	"github.com/iskaa02/taskkit-server/db/models"
)

type (
	rawTask models.Task
	theme   struct {
		Primary   string `json:"primary"`
		Secondary string `json:"secondary"`
	}

	rawList struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Theme theme  `json:"theme"`
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

func (t theme) findTheme(q *models.Queries) (themeID int32, err error) {
	themeArgs := models.FindThemeParams{
		Primary: t.Primary, Secondary: t.Secondary,
	}
	themeID, err = q.FindTheme(context.Background(), themeArgs)
	if err != nil {
		if err == sql.ErrNoRows {
			themeID, err = q.CreateTheme(context.Background(), models.CreateThemeParams(themeArgs))
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	return
}

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
	updateArgs := models.UpdateTaskParams{
		ID:          t.ID,
		Name:        t.Name,
		Subtasks:    t.Subtasks,
		ListID:      t.ListID,
		Description: t.Description,
		Reminder:    t.Reminder,
		Repeat:      t.Repeat,
		IsCompleted: t.IsCompleted,
	}
	updatedRows, err := q.UpdateTask(context.Background(), updateArgs)
	/*
		    if 0 rows have been updated this mean that the column with provided
			id don't exist so we have to insert it
	*/
	if updatedRows == 0 && err == nil {
		err := q.CreateTask(context.Background(), models.CreateTaskParams(updateArgs))
		return err
	}
	return err
}

func (l rawList) Create(q *models.Queries) error {
	themeID, err := l.Theme.findTheme(q)
	err = q.CreateList(context.Background(), models.CreateListParams{
		ID:      l.ID,
		Name:    l.Name,
		ThemeID: themeID,
	})
	return err
}

func (l rawList) Update(q *models.Queries) error {
	themeID, err := l.Theme.findTheme(q)
	if err != nil {
		return err
	}
	updateArgs := models.UpdateListParams{
		ID:      l.ID,
		Name:    l.Name,
		ThemeID: themeID,
	}
	updatedRows, err := q.UpdateList(context.Background(), updateArgs)
	if updatedRows == 0 && err == nil {
		err := q.CreateList(context.Background(), models.CreateListParams(updateArgs))
		return err
	}
	return err
}
