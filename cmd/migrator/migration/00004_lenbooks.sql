-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."lendbooks" (
  "id" uuid   NOT NULL,
  "user_id" uuid NOT NULL,
  "book_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "from" timestamptz DEFAULT now(),
  "to" timestamptz,
  CONSTRAINT "lendbooks_pkey" PRIMARY KEY("id"),
  CONSTRAINT "lendbooks_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE,
  CONSTRAINT "lendbooks_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) NOT DEFERRABLE
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."lendbooks";