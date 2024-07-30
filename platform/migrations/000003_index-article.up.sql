-- Indexes for the Articles Table
CREATE INDEX idx_articles_author ON articles (author);
CREATE INDEX idx_articles_title ON articles (title);
CREATE INDEX idx_articles_deleted_at ON articles (deleted_at);