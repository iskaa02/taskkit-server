package sync

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iskaa02/taskkit-server/db/models"
	"github.com/iskaa02/taskkit-server/ent"
	"github.com/iskaa02/taskkit-server/ent/theme"
	"gopkg.in/guregu/null.v4"
)

type (
	rawTask models.GetNewlyCreatedTasksRow
	Theme   struct {
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
		theme.Or(theme.Secondary(t.Secondary.String), theme.SecondaryIsNil()),
	).OnlyID(ctx)
	if err == sql.ErrNoRows {
		new := a.Create().SetPrimary(t.Primary)
		if t.Secondary.Valid {
			new.SetSecondary(t.Secondary.String)
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

	update := task.Update().
		SetName(t.Name).SetListID(t.ListID).SetSubtasks(&t.Subtasks).SetIsCompleted(t.IsCompleted)
	if t.Description.Valid {
		update.SetNillableDescription(&t.Description.String)
	}
	if t.Reminder.Valid {
		update.SetNillableReminder(&t.Reminder.Time)
	}
	if t.Repeat.Valid {
		update.SetNillableRepeat(&t.Repeat.String)
	}

	return update.Exec(ctx)
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
