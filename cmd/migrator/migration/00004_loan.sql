-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."loans" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "to" timestamptz,
  "deleted_at" timestamptz,
  "book_id" uuid references books,
  "user_id" uuid references users,
  CONSTRAINT "loans_pkey" PRIMARY KEY ("id")
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."loans"
