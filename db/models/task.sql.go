// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: task.sql

package models

import (
	"context"
	"database/sql"

	"github.com/tabbed/pqtype"
)

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
	ID          string
	Name        string
	Subtasks    pqtype.NullRawMessage
	ListID      string
	Description sql.NullString
	Reminder    sql.NullTime
	Repeat      RepeatEnum
	IsCompleted bool
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
DELETE FROM task WHERE id=$1
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, id)
	return err
}

const updateTask = `-- name: UpdateTask :exec
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
	ID          string
	Name        string
	Subtasks    pqtype.NullRawMessage
	ListID      string
	Description sql.NullString
	Reminder    sql.NullTime
	Repeat      RepeatEnum
	IsCompleted bool
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.exec(ctx, q.updateTaskStmt, updateTask,
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
