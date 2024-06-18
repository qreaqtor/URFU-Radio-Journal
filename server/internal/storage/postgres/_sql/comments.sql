DROP TABLE IF EXISTS "comments";
CREATE TABLE "comments" (
    "comment_id" integer NOT NULL PRIMARY KEY,
    "content_ru" text NOT NULL,
    "content_en" text NOT NULL,
    "article" integer NOT NULL,
    "is_approved" boolean NOT NULL,
    "author" character varying(1000) NOT NULL,
    "date_create" date NOT NULL
);

ALTER TABLE "comments" ALTER COLUMN "comment_id" ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "comment_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);