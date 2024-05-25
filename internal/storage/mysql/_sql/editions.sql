DROP TABLE IF EXISTS `editions`;
CREATE TABLE `editions` (
    'edition_id' integer NOT NULL,
    'year' integer NOT NULL,
    'number' integer NOT NULL,
    'volume' integer NOT NULL,
    'cover_path' character varying(1000),
    'file_path' character varying(1000),
    'date' date NOT NULL
)
