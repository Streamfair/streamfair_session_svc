-- name: CreateSession :one
INSERT INTO "session_svc"."Sessions" (
 id,
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
WHERE id = $1 LIMIT 1;

-- name: VerifySession :one
SELECT * FROM "session_svc"."Sessions"
WHERE id = $1 AND is_blocked = false AND expires_at > NOW() LIMIT 1;

-- name: InvalidateSession :exec
UPDATE "session_svc"."Sessions" SET is_blocked = true WHERE id = $1;

-- name: ExtendSession :one
UPDATE "session_svc"."Sessions" SET expires_at = $2 WHERE id = $1 RETURNING *;
