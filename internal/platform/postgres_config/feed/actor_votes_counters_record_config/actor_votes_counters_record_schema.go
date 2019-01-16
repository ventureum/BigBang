package actor_votes_counters_record_config

const TABLE_SCHEMA_FOR_ACTOR_VOTES_COUNTERS_RECORD = `
CREATE TABLE actor_votes_counters_records (
    post_hash TEXT NOT NULL REFERENCES posts(post_hash),
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    latest_reputation BIGINT NOT NULL DEFAULT 0,
    latest_vote_type vote_type_enum NOT NULL,
    latest_reputation_for_upvote BIGINT NOT NULL DEFAULT 0,
    latest_reputation_for_downvote BIGINT NOT NULL DEFAULT 0,
    downvote_count BIGINT NOT NULL DEFAULT 0 check(downvote_count <= 1),
    upvote_count BIGINT NOT NULL DEFAULT 0 check(upvote_count <= 1),
    total_vote_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash, actor)
);
CREATE INDEX actor_votes_counters_records_index ON actor_votes_counters_records(post_hash, actor, latest_reputation, latest_vote_type, latest_reputation_for_upvote, latest_reputation_for_downvote, total_vote_count, updated_at);
`

const TABLE_NAME_FOR_ACTOR_VOTES_COUNTERS_RECORD = "actor_votes_counters_records"
