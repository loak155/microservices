CREATE TABLE "articles" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "url" text NOT NULL,
  "bookmark_count" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);