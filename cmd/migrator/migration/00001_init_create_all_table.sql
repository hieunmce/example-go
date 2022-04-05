-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text,
  "email" text unique ,
  "password" text,
  CONSTRAINT "users_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
