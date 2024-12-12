CREATE TABLE IF NOT EXISTS  "users" (
  "user_id" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT  timestamp '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "files_owned" varchar[] 
);

CREATE TABLE  IF NOT EXISTS  "products" (
    "id" varchar PRIMARY KEY,
    "user_id" varchar NOT NULL,
    "product_name" varchar NOT NULL ,
    "product_description" varchar NOT NULL,
    "product_price" NUMERIC(10,2) NOT NULL,
    "product_urls" varchar[] NOT NULL,
    "compressed_product_images_urls" varchar[] ,
    "created_at" timestamptz NOT NULL DEFAULT (now())

);
CREATE INDEX ON "products" ("user_id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

