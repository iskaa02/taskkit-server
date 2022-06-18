package sync

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/iskaa02/taskkit-server/ent"
	"github.com/iskaa02/taskkit-server/ent/theme"
	"github.com/iskaa02/taskkit-server/types"
	"gopkg.in/guregu/null.v4"
)

type (
	rawTask struct {
		ID          string         `json:"id"`
		Description null.String    `json:"description"`
		Name        string         `json:"name"`
		Subtasks    types.Subtasks `json:"subtasks"`
		Repeat      null.String    `json:"repeat"`
		Reminder    null.Time      `json:"reminder"`
		ListID      string         `json:"list_id"`
		IsCompleted bool           `json:"is_completed"`
	}
	Theme struct {
		Primary   string      `json:"primary"`
		Secondary null.String `json:"secondary"`
	}

	rawList struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Theme Theme  `json:"theme"`
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

func (t Theme) findTheme(c *ent.Tx, ctx context.Context) (themeID int64, err error) {
	a := ent.ThemeClient{}
	themeID, err = c.Theme.Query().Where(
		theme.Primary(t.Primary),
		theme.Or(theme.Secondary(t.Secondary)),
	).OnlyID(ctx)
	if err == sql.ErrNoRows {
		new := a.Create().SetPrimary(t.Primary)
		if t.Secondary.Valid {
			new.SetSecondary(t.Secondary)
		}
		var newTheme *ent.Theme
		newTheme, err = new.Save(ctx)
		themeID = newTheme.ID
	}
	return
}

func (t rawTask) Update(c *ent.Tx, ctx context.Context) error {
	task, err := c.Task.Get(ctx, t.ID)
	// task.Update().Mutation().ListIDs
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = c.Task.Create().SetName(t.Name).
				SetID(t.ID).
				SetIsCompleted(t.IsCompleted).Save(ctx)
		}
		return err
	}
	if task.IsDeleted {
		return errors.New("conflict")
	}

	fmt.Println("listid", t.ListID)
	return task.Update().
		SetName(t.Name).
		SetListID(t.ListID).
		SetSubtasks(&t.Subtasks).
		SetIsCompleted(t.IsCompleted).
		SetDescription(t.Description).
		SetReminder(t.Reminder).
		SetRepeat(t.Repeat).Exec(ctx)
}

func (l rawList) Update(client *ent.Tx, ctx context.Context) error {
	list, err := client.List.Get(ctx, l.ID)
	themeID, err := l.Theme.findTheme(client, ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			err = client.List.Create().SetName(l.ID).SetThemeID(themeID).Exec(ctx)
		}
		return err
	}
	if list.IsDeleted {
		return errors.New("conflict")
	}
	return list.Update().SetThemeID(themeID).SetName(l.Name).Exec(ctx)
}
