ALTER TABLE "articles"
ADD CONSTRAINT "fk_article_edition_id"
FOREIGN KEY ("edition_id")
REFERENCES "editions" ("edition_id")
ON DELETE CASCADE;

ALTER TABLE "comments"
ADD CONSTRAINT "fk_comment_article_id"
FOREIGN KEY ("article")
REFERENCES "articles" ("article_id")
ON DELETE CASCADE;