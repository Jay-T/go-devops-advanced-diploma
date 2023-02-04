CREATE TABLE "account" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "passhash" varchar NOT NULL,
  "blocked" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "secrets" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "key" varchar NOT NULL,
  "value" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "files" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "filename" varchar NOT NULL,
  "filepath" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "metadata" (
  "id" BIGSERIAL PRIMARY KEY,
  "obj_id" bigint NOT NULL,
  "obj_type" varchar NOT NULL,
  "account_id" bigint NOT NULL,
  "key" varchar NOT NULL,
  "value" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "account" ("username");

CREATE INDEX ON "secrets" ("key");

CREATE UNIQUE INDEX ON "secrets" ("account_id", "key");

CREATE INDEX ON "files" ("filename");

CREATE INDEX ON "files" ("filepath");

CREATE UNIQUE INDEX ON "files" ("account_id", "filepath", "filename");

ALTER TABLE "secrets" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "files" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "metadata" ADD FOREIGN KEY ("obj_id") REFERENCES "files" ("id");

ALTER TABLE "metadata" ADD FOREIGN KEY ("obj_id") REFERENCES "secrets" ("id");
