CREATE TABLE "learning_items" (
  "id" text PRIMARY KEY NOT NULL,
  "image_link" text NOT NULL,
  "english_word" text NOT NULL,
  "vietnamese_word" text,
  "english_sentences" text[],
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "completed_at" timestamptz,
  "deleted_at" timestamptz,
  "user_id" text
);

CREATE TABLE "users" (
  "id" text PRIMARY KEY NOT NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "email" text UNIQUE NOT NULL,
  "encrypted_password" text NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

ALTER TABLE "learning_items" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");