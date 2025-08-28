-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE -- PROVIDING "NO KEY" TO PREVENT DEAD LOCK WHILE DB TRANSACTION
; 

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1 -- NUMBER OF ROWS TO RETURN
OFFSET $2 -- NUMBER OF ROWS TO SKIP
; 

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE id = $1;