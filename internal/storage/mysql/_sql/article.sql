DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
    'article_id' integer NOT NULL,
    'title_ru' character varying(1000) NOT NULL,
    'title_en' character varying(1000) NOT NULL,
    'reference_ru' character varying(1000) NOT NULL,
    'reference_en' character varying(1000) NOT NULL,
    'content_ru' text NOT NULL,
    'content_en' text NOT NULL,
    'keywords_ru' text[] NOT NULL,
    'keywords_en' text[] NOT NULL,
    'literature' text[] NOT NULL,
    'edition_id' integer NOT NULL,
    'file_path' character varying(1000) NOT NULL,
    'video_path' character varying(1000) NOT NULL,
    'date_receipt' date NOT NULL,
    'date_acceptance' date NOT NULL,
    'doi' character varying(1000) NOT NULL
)
