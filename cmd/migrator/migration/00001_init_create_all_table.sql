-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."accounts" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
 
  "user_name" varchar(200),
  "digest_password" varchar,
  "type" varchar(200)
);

CREATE TABLE "public"."shops" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  
  "name" varchar(200),
  "address" varchar
);

CREATE TABLE "public"."drinks" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,

  "name" varchar(200),
  "description" varchar,
  "price" SMALLINT
);

CREATE TABLE "public"."details" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,

  "quantity" SMALLINT,
  "drink_id" uuid REFERENCES drinks(id)
);

CREATE TABLE "public"."orders" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,

 "status" varchar(50),
 "order_time" timestamptz,
 "receive_time" timestamptz,
 "shop_id" uuid REFERENCES shops(id),
 "account_id" uuid REFERENCES accounts(id),
 "detail_id" uuid REFERENCES details(id)
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;