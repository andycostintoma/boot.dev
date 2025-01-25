-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, expires_at, revoked_at, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetToken :one
SELECT *
FROM refresh_tokens
WHERE token = $1;

-- name: RevokeToken :one
UPDATE refresh_tokens
SET updated_at = NOW(),
    revoked_at = NOW()
WHERE token = $1
RETURNING *;