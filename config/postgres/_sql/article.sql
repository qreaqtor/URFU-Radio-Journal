DROP TABLE IF EXISTS "articles";
CREATE TABLE "articles" (
    "article_id" integer NOT NULL PRIMARY KEY,
    "title" jsonb NOT NULL,
    "reference" jsonb NOT NULL,
    "content" jsonb NOT NULL,
    "keywords" jsonb NOT NULL,
    "literature" text[] NOT NULL,
    "edition_id" integer NOT NULL,
    "file_path" character varying(1000) NOT NULL,
    "video_path" character varying(1000),
    "date_receipt" date NOT NULL,
    "date_acceptance" date NOT NULL,
    "doi" character varying(1000) NOT NULL,
    "authors" jsonb not null
);

ALTER TABLE "articles" ALTER COLUMN "article_id" ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "article_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE INDEX idx_title_Ru ON "articles" USING gin (to_tsvector('russian', title->>'Ru'));
CREATE INDEX idx_title_Eng ON "articles" USING gin (to_tsvector('english', title->>'Eng'));

CREATE INDEX idx_content_Ru ON "articles" USING gin (to_tsvector('russian', content->>'Ru'));
CREATE INDEX idx_content_Eng ON "articles" USING gin (to_tsvector('english', content->>'Eng'));

CREATE INDEX idx_authors_fullname_ru ON "articles" USING gin (to_tsvector('russian', authors));
CREATE INDEX idx_authors_fullname_eng ON "articles" USING gin (to_tsvector('english', authors));

CREATE INDEX idx_keywords_ru ON "articles" USING gin (to_tsvector('russian', keywords));
CREATE INDEX idx_keywords_eng ON "articles" USING gin (to_tsvector('english', keywords));
