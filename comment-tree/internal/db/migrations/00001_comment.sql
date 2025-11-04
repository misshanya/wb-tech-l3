-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    parent_id UUID REFERENCES comment(id),
    path TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
CREATE INDEX idx_content_fulltext ON comment USING GIN(
    to_tsvector('russian', content), to_tsvector('english', content)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comment;
-- +goose StatementEnd
