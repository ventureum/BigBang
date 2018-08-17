package post_votes_record_config


const TABLE_SCHEMA_FOR_POST_VOTES_RECORD = `
CREATE TABLE post_votes_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    actor TEXT NOT NULL,
    post_hash TEXT NOT NULL,
    vote_type vote_type_enum NOT NULL,
    signed_reputations BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
CREATE INDEX post_votes_records_index ON post_votes_records (actor, post_hash, vote_type, signed_reputations, created_at);
`

const TABLE_NAME_FOR_POST_VOTES_RECORD = "post_votes_records"
