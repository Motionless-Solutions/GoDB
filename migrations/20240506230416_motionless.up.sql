CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "username" varchar NOT NULL,
  "passwordhash" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "createdate" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "role" integer NOT NULL
);
