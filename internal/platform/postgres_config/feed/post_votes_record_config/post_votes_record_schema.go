package post_votes_record_config

const TABLE_SCHEMA_FOR_POST_VOTES_RECORD = `
CREATE TABLE post_votes_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    post_hash TEXT NOT NULL REFERENCES posts(post_hash),
    post_type TEXT NOT NULL,
    vote_type vote_type_enum NOT NULL,
    delta_fuel BIGINT NOT NULL DEFAULT 0,
    delta_reputation BIGINT NOT NULL DEFAULT 0,
    delta_milestone_points BIGINT NOT NULL DEFAULT 0,
    signed_reputation BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
CREATE INDEX post_votes_records_index ON post_votes_records (uuid, actor, post_hash, post_type, vote_type, delta_fuel, delta_reputation, delta_milestone_points, signed_reputation, created_at);
`

const TABLE_NAME_FOR_POST_VOTES_RECORD = "post_votes_records"
