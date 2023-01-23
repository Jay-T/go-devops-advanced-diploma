CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "secrets" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "key" varchar NOT NULL,
  "value" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "files" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "filename" varchar NOT NULL,
  "filepath" varchar NOT NULL,
  "ready" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "secrets" ("key");

CREATE INDEX ON "files" ("filename");

CREATE INDEX ON "files" ("filepath");

CREATE INDEX ON "files" ("filepath", "filename");

COMMENT ON COLUMN "files"."ready" IS 'file ready or not for listing';

ALTER TABLE "secrets" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "files" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");
