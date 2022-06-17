package utils

import (
	"context"
	"database/sql"

	"github.com/iskaa02/taskkit-server/ent"
	"github.com/iskaa02/taskkit-server/ent/theme"
	"gopkg.in/guregu/null.v4"
)

func GetTheme(ctx context.Context, primary string, secondary null.String) (themeID int64, err error) {
	a := ent.ThemeClient{}
	themeID, err = a.Query().Where(
		theme.Primary(primary),
		theme.Or(theme.Secondary(secondary.String), theme.SecondaryIsNil()),
	).OnlyID(ctx)
	if err == sql.ErrNoRows {
		new := a.Create().SetPrimary(primary)
		if secondary.Valid {
			new.SetSecondary(secondary.String)
		}
		var newTheme *ent.Theme
		newTheme, err = new.Save(ctx)
		themeID = newTheme.ID
	}
	return
}
