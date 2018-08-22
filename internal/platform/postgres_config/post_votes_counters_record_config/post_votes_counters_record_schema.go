package post_votes_counters_record_config


const TABLE_SCHEMA_FOR_POST_VOTES_COUNTERS_RECORDS = `
CREATE TABLE post_votes_counters_records (
    post_hash TEXT NOT NULL,
    latest_vote_type vote_type_enum NOT NULL,
    downvote_count BIGINT NOT NULL DEFAULT 0,
    upvote_count BIGINT NOT NULL DEFAULT 0,
    total_vote_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash)
);
CREATE INDEX post_votes_counters_records_index ON post_votes_counters_records (post_hash, latest_vote_type, downvote_count, upvote_count, total_vote_count);
`


const TABLE_NAME_FOR_POST_VOTES_COUNTERS_RECORD = "post_votes_counters_records"
