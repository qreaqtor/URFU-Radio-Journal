DROP TABLE IF EXISTS "redaction";
CREATE TABLE "redaction" (
    "member_id" integer NOT NULL,
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
