-- name: CreateList :exec
INSERT INTO list (id,name,theme_id) VALUES($1,$2,$3);

-- name: UpdateList :exec
UPDATE list SET id=$1, name=$2, theme_id=$3 
    WHERE id=$1;

-- name: DeleteList :exec
DELETE FROM list WHERE id=$1;

-- name: CreateTheme :exec
INSERT INTO theme("primary",secondary)
    VALUES($1,$2);

-- name: FindTheme :one
SELECT id FROM theme WHERE "primary"=$1 AND secondary=$2;
