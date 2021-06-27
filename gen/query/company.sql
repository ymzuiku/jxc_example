-- name: InsertCompany :one
INSERT INTO company (
  account_id,
  name,
  people
) VALUES ($1, $2, $3) RETURNING *;

-- name: SelectCompany :one
SELECT id, name FROM company WHERE account_id = $1 LIMIT 1;


-- name: DeleteAccount :exec
DELETE FROM company WHERE account_id = $1 AND name = $2;