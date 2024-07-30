CREATE TABLE articles (
    id UUID PRIMARY KEY,
    author VARCHAR(255),
    title VARCHAR(255),
    body TEXT,
    created_by UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

-- Indexes for the Articles Table
CREATE INDEX idx_articles_author ON articles (author);
CREATE INDEX idx_articles_title ON articles (title);
CREATE INDEX idx_articles_deleted_at ON articles (deleted_at);
