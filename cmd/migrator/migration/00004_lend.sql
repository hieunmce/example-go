-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."lends" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "book_id" uuid NOT NULL,
  "name" text,
  "from" timestamptz,
  "to" timestamptz,
  CONSTRAINT "lends_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "lends_book_id_fkey" FOREIGN KEY ("book_id") REFERENCES books("id") NOT DEFERRABLE,
  CONSTRAINT "lends_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES users("id") NOT DEFERRABLE
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."lends";

