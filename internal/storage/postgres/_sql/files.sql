DROP TABLE IF EXISTS "fileinfo";
CREATE TABLE "fileinfo" (
    "file_id" integer PRIMARY KEY,
    "filename" character varying(1000),
    "backet" integer NOT NULL
);