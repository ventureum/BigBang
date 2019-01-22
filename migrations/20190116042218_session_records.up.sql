CREATE TABLE session_records (
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    post_hash TEXT NOT NULL REFERENCES posts(post_hash),
    start_time BIGINT NOT NULL,
    end_time BIGINT NOT NULL,
    content JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash)
);
CREATE INDEX session_records_index ON session_records (actor, post_hash);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON session_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();