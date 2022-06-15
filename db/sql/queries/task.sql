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

-- name: CheckTaskExist :one
SELECT EXISTS(SELECT 1 FROM task WHERE ID =$1);

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

