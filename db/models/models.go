// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package models

import (
	"github.com/tabbed/pqtype"
	null "gopkg.in/guregu/null.v4"
)

type List struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ThemeID      int32     `json:"theme_id"`
	LastModified null.Time `json:"last_modified"`
}

type Task struct {
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Subtasks     pqtype.NullRawMessage `json:"subtasks"`
	ListID       string                `json:"list_id"`
	Description  null.String           `json:"description"`
	Reminder     null.Time             `json:"reminder"`
	Repeat       null.String           `json:"repeat"`
	IsCompleted  bool                  `json:"is_completed"`
	LastModified null.Time             `json:"last_modified"`
}

type Theme struct {
	ID        int32  `json:"id"`
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}
