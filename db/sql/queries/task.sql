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
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
ON CONFLICT (id) DO 
    UPDATE SET
        name = excluded.name, 
        subtasks = excluded.subtasks,
        list_id = excluded.list_id,
        description = excluded.description,
        reminder = excluded.reminder,
        repeat = excluded.repeat,
        is_completed = excluded.is_completed;

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
DELETE FROM task WHERE id=$1;
