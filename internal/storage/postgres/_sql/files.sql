DROP TABLE IF EXISTS "files";
CREATE TABLE "files" (
    "file_id" integer PRIMARY KEY,
    "content_type" character varying(1000),
    "filename" character varying(1000),
    "size" integer NOT NULL,
);