-- name: SelectUserByPhone :one
SELECT * FROM users WHERE phone = $1 LIMIT 1;

-- name: SelectUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: InsertUser :exec
INSERT INTO users (
  name,
  phone,
  password
) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users SET password = $2 WHERE id = $1;

-- name: DeleteUserByPhone :exec
DELETE FROM users WHERE phone = $1;