-- name: CreateAccount :one
INSERT INTO account (
  username,
  passhash
) VALUES (
  $1, $2
)
RETURNING *;

-- name: BlockAccount :exec
UPDATE account
  set blocked = true
WHERE username = $1;

-- name: GetAccount :one
SELECT * FROM account
WHERE username = $1 LIMIT 1;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE username = $1;