-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."user_lend_books" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "user_id" uuid,
  "book_id" uuid, 
  "from" timestamptz,
  "to" timestamptz,

  CONSTRAINT "user_lend_books_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "user_lend_books_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE,
  CONSTRAINT "user_lend_books_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) NOT DEFERRABLE
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."user_lend_books";
