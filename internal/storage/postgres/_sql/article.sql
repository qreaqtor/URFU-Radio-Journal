DROP TABLE IF EXISTS "articles";
CREATE TABLE "articles" (
    "article_id" integer NOT NULL PRIMARY KEY,
    "title" json NOT NULL,
    "reference" json NOT NULL,
    "content" json NOT NULL,
    "keywords" json NOT NULL,
    "literature" text[] NOT NULL,
    "edition_id" integer NOT NULL,
    "file_path" character varying(1000) NOT NULL,
    "video_path" character varying(1000) NOT NULL,
    "date_receipt" date NOT NULL,
    "date_acceptance" date NOT NULL,
    "doi" character varying(1000) NOT NULL,
    "authors" json not null
);

ALTER TABLE "articles" ALTER COLUMN "article_id" ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "article_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);