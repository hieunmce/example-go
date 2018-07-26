-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "public"."lendingabook" (
    "id" uuid NOT NULL,
    "created_at" timestamptz DEFAULT now(),
    "deleted_at" timestamptz,
    "book_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "from" timestamptz DEFAULT now(),
    "to" timestamptz,
    CONSTRAINT "lendingabook_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "lendingabook_book_id_fkey" FOREIGN KEY (book_id) REFERENCES books(id) NOT DEFERRABLE,
    CONSTRAINT "lendingabook_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) NOT DEFERRABLE
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "public"."lendingabook";
