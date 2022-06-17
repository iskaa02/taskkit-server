package sync

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iskaa02/taskkit-server/db/models"
)

type (
	rawTask models.GetNewlyCreatedTasksRow
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
	isDeleted, err := q.CheckTaskIsDeleted(context.Background(), t.ID)
	if err != nil {
		// Task don't exist
		if err == sql.ErrNoRows {
			return q.CreateTask(context.Background(), models.CreateTaskParams(updateArgs))
		} else {
			return err
		}
	}

	if isDeleted {
		return errors.New("conflict")
	}
	_, err = q.UpdateTask(context.Background(), updateArgs)
	return err
}

func (l rawList) Update(q *models.Queries) error {
	// exist, err := q.CheckListExist(context.Background(), l.ID)
	isDeleted, err := q.CheckListIsDeleted(context.Background(), l.ID)
	if err != nil {
		return err
	}
	themeID, err := l.Theme.findTheme(q)
	if err != nil {
		return err
	}
	updateArgs := models.UpdateListParams{
		ID:      l.ID,
		Name:    l.Name,
		ThemeID: themeID,
	}
	if err != nil {
		// Task don't exist
		if err == sql.ErrNoRows {
			return q.CreateList(context.Background(), models.CreateListParams(updateArgs))
		} else {
			return err
		}
	}
	if isDeleted {
		return errors.New("conflict")
	}
	_, err = q.UpdateList(context.Background(), updateArgs)
	return err
}
