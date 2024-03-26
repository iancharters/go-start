-- name: CreateUser :one
INSERT INTO users(username, PASSWORD)
  VALUES ($1, $2)
RETURNING
  *;

-- name: GetUserByID :one
SELECT
  *
FROM
  users
WHERE
  users.id = $1;

-- name: GetUserByUsername :one
SELECT
  *
FROM
  users
WHERE
  users.username = $1;

-- name: ListUsers :many
SELECT
  *
FROM
  users
ORDER BY
  id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE users.id = $1;

