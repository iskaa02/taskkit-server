-- name: CreateTask :exec
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
VALUES($1,$2,$3,$4,$5,$6,$7,$8);

-- name: CheckTaskIsDeleted :one
SELECT is_deleted
FROM task WHERE id =$1;

-- name: UpdateTask :execrows
UPDATE task SET
    name=$2,
    subtasks=$3,
    list_id=$4,
    description=$5,
    reminder=$6,
    repeat=$7,
    is_completed=$8
WHERE id=$1;

-- name: DeleteTask :exec
UPDATE task SET is_deleted=true WHERE id=$1;

-- name: GetNewlyCreatedTasks :many
SELECT id,name,subtasks,list_id,description,reminder,repeat,is_completed,is_deleted FROM task
WHERE created_at >= $1;

-- name: GetNewlyUpdatedTasks :many
SELECT id,name,subtasks,list_id,description,reminder,repeat,is_completed,is_deleted FROM task
WHERE created_at <= $1 AND last_modified >= $1;

-- name: GetNewlyDeletedTasks :many
SELECT id FROM task
WHERE is_deleted=TRUE AND last_modified >= $1;
