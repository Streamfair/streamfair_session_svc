#!/bin/sh

set -e

echo "Running DB migrations"
/streamfair_session_svc/migrate -path /streamfair_session_svc/migration -database "$DB_SOURCE_SESSION_SERVICE" -verbose up

echo "Starting Session Service"
exec "$@"