-- name: CreateList :exec
INSERT INTO list (id,name,theme_id) VALUES($1,$2,$3);

-- name: UpdateList :execrows
UPDATE list SET id=$1, name=$2, theme_id=$3 
    WHERE id=$1;

-- name: DeleteList :exec
UPDATE list SET is_deleted=true WHERE id=$1;

-- name: CheckListExist :one
SELECT EXISTS(SELECT 1 FROM list WHERE id =$1);

-- name: CreateTheme :one
INSERT INTO theme("primary",secondary)
    VALUES($1,$2) RETURNING id;

-- name: FindTheme :one
SELECT id FROM theme WHERE "primary"=$1 AND secondary=$2;
