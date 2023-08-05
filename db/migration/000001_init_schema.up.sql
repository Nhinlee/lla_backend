CREATE TABLE learning_item (
    id text primary key NOT NULL,
    image_link text NOT NULL,
    english_word text NOT NULL,
    vietnamese_word text NOT NULL,
    english_sentences text[]
);

CREATE TABLE "users" (
  user_id text PRIMARY KEY NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL UNIQUE,
  "password" text NOT NULL
);