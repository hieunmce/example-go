-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."books" (
  "id" uuid NOT NULL,
  "category_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "name" text,
  "author" text,
  "description" text,
  CONSTRAINT "books_pkey" PRIMARY KEY ("id","category_id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."books";