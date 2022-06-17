// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: task.sql

package models

import (
	"context"
	"time"

	"github.com/tabbed/pqtype"
	null "gopkg.in/guregu/null.v4"
)

const checkTaskIsDeleted = `-- name: CheckTaskIsDeleted :one
SELECT is_deleted
FROM task WHERE id =$1
`

func (q *Queries) CheckTaskIsDeleted(ctx context.Context, id string) (bool, error) {
	row := q.queryRow(ctx, q.checkTaskIsDeletedStmt, checkTaskIsDeleted, id)
	var is_deleted bool
	err := row.Scan(&is_deleted)
	return is_deleted, err
}

const createTask = `-- name: CreateTask :exec
INSERT INTO task (
    id,
    name,
    subtasks,
    list_id,
    description,
    reminder,
    repeat,
    is_completed
)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
`

type CreateTaskParams struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Subtasks    pqtype.NullRawMessage `json:"subtasks"`
	ListID      string                `json:"list_id"`
	Description null.String           `json:"description"`
	Reminder    null.Time             `json:"reminder"`
	Repeat      null.String           `json:"repeat"`
	IsCompleted bool                  `json:"is_completed"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) error {
	_, err := q.exec(ctx, q.createTaskStmt, createTask,
		arg.ID,
		arg.Name,
		arg.Subtasks,
		arg.ListID,
		arg.Description,
		arg.Reminder,
		arg.Repeat,
		arg.IsCompleted,
	)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
UPDATE task SET is_deleted=true WHERE id=$1
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, id)
	return err
}

const getNewlyCreatedTasks = `-- name: GetNewlyCreatedTasks :many
SELECT id,name,subtasks,list_id,description,reminder,repeat,is_completed,is_deleted FROM task
WHERE created_at >= $1
`

type GetNewlyCreatedTasksRow struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Subtasks    pqtype.NullRawMessage `json:"subtasks"`
	ListID      string                `json:"list_id"`
	Description null.String           `json:"description"`
	Reminder    null.Time             `json:"reminder"`
	Repeat      null.String           `json:"repeat"`
	IsCompleted bool                  `json:"is_completed"`
	IsDeleted   bool                  `json:"is_deleted"`
}

func (q *Queries) GetNewlyCreatedTasks(ctx context.Context, createdAt time.Time) ([]GetNewlyCreatedTasksRow, error) {
	rows, err := q.query(ctx, q.getNewlyCreatedTasksStmt, getNewlyCreatedTasks, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNewlyCreatedTasksRow
	for rows.Next() {
		var i GetNewlyCreatedTasksRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Subtasks,
			&i.ListID,
			&i.Description,
			&i.Reminder,
			&i.Repeat,
			&i.IsCompleted,
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

const getNewlyDeletedTasks = `-- name: GetNewlyDeletedTasks :many
SELECT id FROM task
WHERE is_deleted=TRUE AND last_modified >= $1
`

func (q *Queries) GetNewlyDeletedTasks(ctx context.Context, lastModified time.Time) ([]string, error) {
	rows, err := q.query(ctx, q.getNewlyDeletedTasksStmt, getNewlyDeletedTasks, lastModified)
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

const getNewlyUpdatedTasks = `-- name: GetNewlyUpdatedTasks :many
SELECT id,name,subtasks,list_id,description,reminder,repeat,is_completed,is_deleted FROM task
WHERE created_at <= $1 AND last_modified >= $1
`

type GetNewlyUpdatedTasksRow struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Subtasks    pqtype.NullRawMessage `json:"subtasks"`
	ListID      string                `json:"list_id"`
	Description null.String           `json:"description"`
	Reminder    null.Time             `json:"reminder"`
	Repeat      null.String           `json:"repeat"`
	IsCompleted bool                  `json:"is_completed"`
	IsDeleted   bool                  `json:"is_deleted"`
}

func (q *Queries) GetNewlyUpdatedTasks(ctx context.Context, createdAt time.Time) ([]GetNewlyUpdatedTasksRow, error) {
	rows, err := q.query(ctx, q.getNewlyUpdatedTasksStmt, getNewlyUpdatedTasks, createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNewlyUpdatedTasksRow
	for rows.Next() {
		var i GetNewlyUpdatedTasksRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Subtasks,
			&i.ListID,
			&i.Description,
			&i.Reminder,
			&i.Repeat,
			&i.IsCompleted,
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

const updateTask = `-- name: UpdateTask :execrows
UPDATE task SET
    name=$2,
    subtasks=$3,
    list_id=$4,
    description=$5,
    reminder=$6,
    repeat=$7,
    is_completed=$8
WHERE id=$1
`

type UpdateTaskParams struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Subtasks    pqtype.NullRawMessage `json:"subtasks"`
	ListID      string                `json:"list_id"`
	Description null.String           `json:"description"`
	Reminder    null.Time             `json:"reminder"`
	Repeat      null.String           `json:"repeat"`
	IsCompleted bool                  `json:"is_completed"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (int64, error) {
	result, err := q.exec(ctx, q.updateTaskStmt, updateTask,
		arg.ID,
		arg.Name,
		arg.Subtasks,
		arg.ListID,
		arg.Description,
		arg.Reminder,
		arg.Repeat,
		arg.IsCompleted,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
