-- name: CreateList :exec
INSERT INTO list (id,name,theme_id) VALUES($1,$2,$3);

-- name: UpdateList :execrows
UPDATE list SET id=$1, name=$2, theme_id=$3 
    WHERE id=$1;

-- name: DeleteList :exec
UPDATE list SET is_deleted=true WHERE id=$1;

-- name: CheckListIsDeleted :one
SELECT is_deleted
FROM list WHERE id =$1;

-- name: CreateTheme :one
INSERT INTO theme("primary",secondary)
    VALUES($1,$2) RETURNING id;

-- name: FindTheme :one
SELECT id FROM theme WHERE "primary"=$1 AND secondary=$2;

-- name: GetNewlyCreatedLists :many
SELECT l.id,name,t."primary",t.secondary,is_deleted FROM list l
JOIN theme t ON theme_id=t.id 
WHERE created_at >= $1;

-- name: GetNewlyUpdatedLists :many
SELECT l.id,name,t."primary",t.secondary,is_deleted FROM list l
JOIN theme t ON theme_id=t.id 
WHERE created_at <= $1 AND last_modified >= $1;

-- name: GetNewlyDeletedLists :many
SELECT id FROM list
WHERE is_deleted=TRUE AND last_modified >= $1;
