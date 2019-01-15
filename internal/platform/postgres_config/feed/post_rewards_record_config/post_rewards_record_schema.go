package post_rewards_record_config

const TABLE_SCHEMA_FOR_POST_REWARDS_RECORD = `
CREATE TABLE post_rewards_records(
    post_hash TEXT NOT NULL REFERENCES posts(post_hash),
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    post_type TEXT NOT NULL,
    object TEXT NOT NULL,
    post_time TIMESTAMPTZ NOT NULL,
    delta_fuel BIGINT NOT NULL DEFAULT 0,
    delta_reputation BIGINT NOT NULL DEFAULT 0,
    delta_milestone_points BIGINT NOT NULL DEFAULT 0,
    withdrawable_mps BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash)
);
CREATE INDEX post_rewards_records_index ON post_rewards_records (post_hash, actor, post_type, delta_fuel, delta_reputation, delta_milestone_points, withdrawable_mps);
`

const TABLE_NAME_FOR_POST_REWARDS_RECORD = "post_rewards_records"
