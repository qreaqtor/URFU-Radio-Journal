DROP TABLE IF EXISTS "redaction";
CREATE TABLE "redaction" (
    "member_id" integer NOT NULL PRIMARY KEY,
    "fullname_ru" character varying(1000) NOT NULL,
    "fullname_en" character varying(1000) NOT NULL,
    "description_ru" character varying(1000) NOT NULL,
    "location_ru" character varying(1000) NOT NULL,
    "email" character varying(1000) NOT NULL,
    "photo_path" character varying(1000) NOT NULL,
    "date_join" date NOT NULL,
    "rank" character varying(1000) NOT NULL,
    "description_en" character varying(1000) NOT NULL,
    "content_ru" character varying(1000) NOT NULL,
    "content_en" character varying(1000) NOT NULL,
    "location_en" character varying(1000) NOT NULL
);

ALTER TABLE "redaction" ALTER COLUMN "member_id" ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME "redaction_member_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);