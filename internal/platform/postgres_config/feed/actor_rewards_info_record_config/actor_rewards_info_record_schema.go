package actor_rewards_info_record_config

const TABLE_SCHEMA_FOR_ACTOR_REWARDS_INFO_RECORD = `
CREATE TABLE actor_rewards_info_records (
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    fuel BIGINT NOT NULL DEFAULT 0,
    reputation BIGINT NOT NULL DEFAULT 0,
    milestone_points_from_votes BIGINT NOT NULL DEFAULT 0,
    milestone_points_from_posts BIGINT NOT NULL DEFAULT 0,
    milestone_points_from_others BIGINT NOT NULL DEFAULT 0,
    consumed_milestone_points BIGINT NOT NULL DEFAULT 0,
    milestone_points BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
CREATE INDEX actor_rewards_info_records_index ON actor_rewards_info_records (actor, fuel, reputation, milestone_points);
`

const TABLE_NAME_FOR_ACTOR_REWARDS_INFO_RECORD = "actor_rewards_info_records"
