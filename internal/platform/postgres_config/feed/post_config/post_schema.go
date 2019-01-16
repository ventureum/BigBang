package post_config

const TABLE_SCHEMA_FOR_POST = `
CREATE TABLE posts (
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    board_id TEXT NOT NULL,
    parent_hash TEXT NOT NULL,
    post_hash TEXT NOT NULL,
    post_type TEXT NOT NULL,
    content JSONB NOT NULL,
    update_count BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash)
);
CREATE INDEX posts_index ON posts (post_hash, update_count);
`

const TABLE_NAME_FOR_POST = "posts"
