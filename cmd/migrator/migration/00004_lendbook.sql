-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."lendbook" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "user_id" uuid NOT NULL,
  "book_id" uuid NOT NULL,
  "from" timestamptz ,
  "to" timestamptz,

  CONSTRAINT "lendbook" PRIMARY KEY ("id")
  CONSTRAINT "lendbook_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE
  CONSTRAINT "lendbook_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) NOT DEFERRABLE

) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."lendbook";