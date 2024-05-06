-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  passwordhash,
  fullname,
  createdate,
  role
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    username = $3,
    passwordhash = $4,
    fullname = $5,
    createdate = $6,
    role = $7
WHERE id = $1
RETURNING *;
