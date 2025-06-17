-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS renditions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        video_id TEXT NOT NULL REFERENCES videos (id) ON DELETE CASCADE,
        label TEXT NOT NULL, -- e.g. '240p', '360p', '720p'
        bitrate INTEGER NOT NULL, -- bits per second, e.g. 300000
        width INTEGER NOT NULL, -- e.g. 426
        height INTEGER NOT NULL, -- e.g. 240
        path TEXT NOT NULL, -- filesystem path relative to your /videos/{id} folder
        UNIQUE (video_id, label) -- prevent duplicate renditions
    );

-- Optional: index on video_id for faster joins
CREATE INDEX IF NOT EXISTS idx_renditions_video ON renditions (video_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS renditions;

-- +goose StatementEnd