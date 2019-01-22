CREATE TABLE post_replies_records (
    post_hash TEXT NOT NULL REFERENCES posts(post_hash),
    reply_hash TEXT UNIQUE NOT NULL REFERENCES posts(post_hash),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash, reply_hash)
);
CREATE INDEX post_replies_records_index ON post_replies_records (post_hash, reply_hash);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON post_replies_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
