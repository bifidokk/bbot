-- Create "chat" table
CREATE TABLE "chat" (
  "id" bigserial NOT NULL,
  "telegram_id" bigint NOT NULL,
  "title" character varying(255) NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
