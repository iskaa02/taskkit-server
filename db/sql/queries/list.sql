-- name: CreateList :exec
-- if id don't exist insert else update
INSERT INTO list (id,name,theme_id) 
VALUES ($1,$2,$3)
ON CONFLICT (id) DO 
    UPDATE SET
        name = excluded.name, 
        theme_id = excluded.theme_id;

-- name: UpdateList :execrows
UPDATE list SET id=$1, name=$2, theme_id=$3 
    WHERE id=$1;

-- name: DeleteList :exec
DELETE FROM list WHERE id=$1;


-- name: CreateTheme :one
INSERT INTO theme("primary",secondary)
    VALUES($1,$2) RETURNING id;

-- name: FindTheme :one
SELECT id FROM theme WHERE "primary"=$1 AND secondary=$2;
