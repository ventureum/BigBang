package post_reputations_record_config


const TABLE_SCHEMA_FOR_POST_REPUTATIONS_RECORD = `
CREATE TABLE post_reputations_records (
    post_hash TEXT NOT NULL,
    actor TEXT NOT NULL,
    reputations BIGINT NOT NULL DEFAULT 0 check(reputations >= 0),
    latest_vote_type vote_type_enum NOT NULL,
    total_vote_count BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash, actor)
);
CREATE INDEX post_reputations_records_index ON post_reputations_records (post_hash, actor, reputations, latest_vote_type, total_vote_count,  updated_at);
`

const TABLE_NAME_FOR_POST_REPUTATIONS_RECORD = "post_reputations_records"
