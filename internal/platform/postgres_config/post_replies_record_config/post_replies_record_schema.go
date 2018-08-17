package post_replies_record_config

const TABLE_SCHEMA_FOR_POST_REPLIES_RECORD = `
CREATE TABLE post_replies_records (
    post_hash TEXT NOT NULL,
    reply_hash TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash, reply_hash)
);
CREATE INDEX post_replies_records_index ON post_replies_records (post_hash, reply_hash);
`

const TABLE_NAME_FOR_POST_REPLIES_RECORD = "post_replies_records"
