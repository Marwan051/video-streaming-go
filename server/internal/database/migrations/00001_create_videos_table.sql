-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS videos (
        id TEXT PRIMARY KEY, -- UUID or other unique ID
        filename TEXT NOT NULL, -- original filename
        upload_at TEXT NOT NULL DEFAULT (DATETIME ('now')), -- ISO‑8601 timestamp
        duration INTEGER, -- duration in seconds
        width INTEGER, -- original width
        height INTEGER, -- original height
        status TEXT NOT NULL DEFAULT 'uploaded' -- e.g. uploaded|processing|ready
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS videos;

-- +goose StatementEnd