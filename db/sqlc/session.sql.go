// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
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
RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (SessionSvcSession, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	var i SessionSvcSession
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const extendSession = `-- name: ExtendSession :one
UPDATE "session_svc"."Sessions" SET expires_at = $2 WHERE id = $1 RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type ExtendSessionParams struct {
	ID        uuid.UUID `json:"id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) ExtendSession(ctx context.Context, arg ExtendSessionParams) (SessionSvcSession, error) {
	row := q.db.QueryRow(ctx, extendSession, arg.ID, arg.ExpiresAt)
	var i SessionSvcSession
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at FROM "session_svc"."Sessions"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (SessionSvcSession, error) {
	row := q.db.QueryRow(ctx, getSession, id)
	var i SessionSvcSession
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const invalidateSession = `-- name: InvalidateSession :exec
UPDATE "session_svc"."Sessions" SET is_blocked = true WHERE id = $1
`

func (q *Queries) InvalidateSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, invalidateSession, id)
	return err
}

const verifySession = `-- name: VerifySession :one
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at FROM "session_svc"."Sessions"
WHERE id = $1 AND is_blocked = false AND expires_at > NOW() LIMIT 1
`

func (q *Queries) VerifySession(ctx context.Context, id uuid.UUID) (SessionSvcSession, error) {
	row := q.db.QueryRow(ctx, verifySession, id)
	var i SessionSvcSession
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
