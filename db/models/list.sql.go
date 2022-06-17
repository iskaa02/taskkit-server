// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: list.sql

package models

import (
	"context"
	"time"
)

const checkListIsDeleted = `-- name: CheckListIsDeleted :one
SELECT is_deleted
FROM list WHERE id =$1
`

func (q *Queries) CheckListIsDeleted(ctx context.Context, id string) (bool, error) {
	row := q.queryRow(ctx, q.checkListIsDeletedStmt, checkListIsDeleted, id)
	var is_deleted bool
	err := row.Scan(&is_deleted)
	return is_deleted, err
}

const createList = `-- name: CreateList :exec
INSERT INTO list (id,name,theme_id) VALUES($1,$2,$3)
`

type CreateListParams struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ThemeID int32  `json:"theme_id"`
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) error {
	_, err := q.exec(ctx, q.createListStmt, createList, arg.ID, arg.Name, arg.ThemeID)
	return err
}

const createTheme = `-- name: CreateTheme :one
INSERT INTO theme("primary",secondary)
    VALUES($1,$2) RETURNING id
`

type CreateThemeParams struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

func (q *Queries) CreateTheme(ctx context.Context, arg CreateThemeParams) (int32, error) {
	row := q.queryRow(ctx, q.createThemeStmt, createTheme, arg.Primary, arg.Secondary)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteList = `-- name: DeleteList :exec
UPDATE list SET is_deleted=true WHERE id=$1
`

func (q *Queries) DeleteList(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteListStmt, deleteList, id)
	return err
}

const findTheme = `-- name: FindTheme :one
SELECT id FROM theme WHERE "primary"=$1 AND secondary=$2
`

type FindThemeParams struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

func (q *Queries) FindTheme(ctx context.Context, arg FindThemeParams) (int32, error) {
	row := q.queryRow(ctx, q.findThemeStmt, findTheme, arg.Primary, arg.Secondary)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getNewlyCreatedLists = `-- name: GetNewlyCreatedLists :many
SELECT l.id,name,t."primary",t.secondary,is_deleted FROM list l
JOIN theme t ON theme_id=t.id 
WHERE created_at >= $1
`

type GetNewlyCreatedListsRow struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	IsDeleted bool   `json:"is_deleted"`
}

func (q *Queries) GetNewlyCreatedLists(ctx context.Context, createdAt time.Time) ([]GetNewlyCreatedListsRow, error) {
	rows, err := q.query(ctx, q.getNewlyCreatedListsStmt, getNewlyCreatedLists, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNewlyCreatedListsRow
	for rows.Next() {
		var i GetNewlyCreatedListsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Primary,
			&i.Secondary,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNewlyDeletedLists = `-- name: GetNewlyDeletedLists :many
SELECT id FROM list
WHERE is_deleted=TRUE AND last_modified >= $1
`

func (q *Queries) GetNewlyDeletedLists(ctx context.Context, lastModified time.Time) ([]string, error) {
	rows, err := q.query(ctx, q.getNewlyDeletedListsStmt, getNewlyDeletedLists, lastModified)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNewlyUpdatedLists = `-- name: GetNewlyUpdatedLists :many
SELECT l.id,name,t."primary",t.secondary,is_deleted FROM list l
JOIN theme t ON theme_id=t.id 
WHERE created_at <= $1 AND last_modified >= $1
`

type GetNewlyUpdatedListsRow struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
	IsDeleted bool   `json:"is_deleted"`
}

func (q *Queries) GetNewlyUpdatedLists(ctx context.Context, createdAt time.Time) ([]GetNewlyUpdatedListsRow, error) {
	rows, err := q.query(ctx, q.getNewlyUpdatedListsStmt, getNewlyUpdatedLists, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNewlyUpdatedListsRow
	for rows.Next() {
		var i GetNewlyUpdatedListsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Primary,
			&i.Secondary,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateList = `-- name: UpdateList :execrows
UPDATE list SET id=$1, name=$2, theme_id=$3 
    WHERE id=$1
`

type UpdateListParams struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ThemeID int32  `json:"theme_id"`
}

func (q *Queries) UpdateList(ctx context.Context, arg UpdateListParams) (int64, error) {
	result, err := q.exec(ctx, q.updateListStmt, updateList, arg.ID, arg.Name, arg.ThemeID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
