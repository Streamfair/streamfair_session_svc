CREATE SCHEMA "session_svc";

CREATE TABLE "session_svc"."Sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_session_username" ON "session_svc"."Sessions" ("username");

CREATE INDEX "idx_user_sessions_created_at" ON "session_svc"."Sessions" ("created_at");