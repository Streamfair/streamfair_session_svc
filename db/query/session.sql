-- name: CreateSession :one
INSERT INTO "session_svc"."Sessions" (
 uuid,
 username,
 refresh_token,
 user_agent,
 client_ip,
 is_blocked,
 expires_at
) VALUES (
 $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetSession :one
SELECT * FROM "session_svc"."Sessions"
WHERE uuid = $1 LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM "session_svc"."Sessions" WHERE uuid = $1;

-- name: RevokeSession :exec
UPDATE "session_svc"."Sessions" SET is_blocked = true WHERE uuid = $1;

-- name: ExtendSession :one
UPDATE "session_svc"."Sessions" SET expires_at = $2 WHERE uuid = $1 RETURNING *;
