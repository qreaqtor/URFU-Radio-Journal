DROP TABLE IF EXISTS "comments";
CREATE TABLE "comments" (
    "comment_id" integer NOT NULL,
    "content_ru" text NOT NULL,
    "content_en" text NOT NULL,
    "article" integer NOT NULL,
    "is_approved" boolean NOT NULL,
    "author" character varying(1000) NOT NULL,
    "date_create" date NOT NULL
);