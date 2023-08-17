CREATE TABLE "topics" (
  "id" text PRIMARY KEY NOT NULL,
  "name" text UNIQUE NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

ALTER TABLE "learning_items" ADD COLUMN topic_id text;

ALTER TABLE "learning_items" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id");