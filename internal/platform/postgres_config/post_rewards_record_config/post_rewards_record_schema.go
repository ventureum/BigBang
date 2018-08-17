package post_rewards_record_config


const TABLE_SCHEMA_FOR_POST_REWARDS_RECORD = `
CREATE TABLE post_rewards_records (
    post_hash TEXT NOT NULL,
    rewards BIGINT NOT NULL DEFAULT 0 check(rewards >= 0),
    latest_median_reputations BIGINT NOT NULL,
    withdrawable_rewards BIGINT NOT NULL DEFAULT 0 check(withdrawable_rewards >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (post_hash)
);
CREATE INDEX post_rewards_records_index ON post_rewards_records (post_hash, rewards);
`

const TABLE_NAME_FOR_POST_REWARDS_RECORD = "post_rewards_records"
