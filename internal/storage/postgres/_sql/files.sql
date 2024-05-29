CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS "fileinfo";
CREATE TABLE "fileinfo" (
    "file_id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "bucket" character varying(1000)
);