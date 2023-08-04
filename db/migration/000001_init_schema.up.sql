CREATE TABLE learning_item (
    id text primary key,
    image_link text,
    english_word text,
    vietnamese_word text,
    english_sentences text[]
);

CREATE TABLE "users" (
  user_id text PRIMARY KEY NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL UNIQUE,
  "password" text NOT NULL
);