-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."books" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "category_id" uuid NOT NULL,
  "name" text,
  "description" text,
  "author" text,
  CONSTRAINT "books_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "books_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES categories("id") NOT DEFERRABLE
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."books";