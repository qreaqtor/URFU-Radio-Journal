DROP TABLE IF EXISTS "authors";
CREATE TABLE "authors" (
    "author_id" integer NOT NULL PRIMARY KEY,
    "fullname_ru" character varying(1000) NOT NULL,
    "fullname_en" character varying(1000) NOT NULL,
    "affiliation" character varying(1000) NOT NULL,
    "email" character varying(1000) NOT NULL,
    "article_id" integer NOT NULL
);

ALTER TABLE "authors" ALTER COLUMN "author_id" ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "author_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);