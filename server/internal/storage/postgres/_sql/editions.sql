DROP TABLE IF EXISTS "editions";
CREATE TABLE "editions" (
    "edition_id" integer NOT NULL PRIMARY KEY,
    "year" integer NOT NULL,
    "number" integer NOT NULL,
    "volume" integer NOT NULL,
    "cover_path" character varying(1000),
    "file_path" character varying(1000),
    "date" date NOT NULL
);

ALTER TABLE "editions" ALTER COLUMN "edition_id" ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "edition_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);