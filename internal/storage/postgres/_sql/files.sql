DROP TABLE IF EXISTS "fileinfo";
CREATE TABLE "fileinfo" (
    "file_id" SERIAL PRIMARY KEY,
    "filename" character varying(1000),
    "bucket" character varying(1000)
);