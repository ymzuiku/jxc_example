-- name: SelectAccountByPhone :one
SELECT * FROM account WHERE phone = $1 LIMIT 1;

-- name: SelectAccountById :one
SELECT * FROM account WHERE id = $1;


-- name: InsertAccount :exec
INSERT INTO account (
  name,
  phone,
  password
) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAccountPassword :exec
UPDATE account SET password = $2 WHERE id = $1;

-- name: DeleteAccountByPhone :exec
DELETE FROM account WHERE phone = $1;