package actor_milestone_points_redeem_history_record_config

const TABLE_SCHEMA_FOR_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD = `
CREATE TABLE actor_milestone_points_redeem_history_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor),
    redeem_block BIGINT NOT NULL DEFAULT 0,
    token_pool BIGINT NOT NULL DEFAULT 0,
    total_enrolled_milestone_points BIGINT NOT NULL DEFAULT 0,
    targeted_milestone_points BIGINT NOT NULL DEFAULT 0,
    actual_milestone_points BIGINT NOT NULL DEFAULT 0,
    consumed_milestone_points BIGINT NOT NULL DEFAULT 0,
    redeemed_tokens BIGINT NOT NULL DEFAULT 0,
    submitted_at TIMESTAMPTZ NOT NULL,
    executed_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
CREATE INDEX actor_milestone_points_redeem_history_records_index ON actor_milestone_points_redeem_history_records (
     actor, redeem_block, redeemed_tokens, submitted_at, executed_at);
`

const TABLE_NAME_FOR_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD = "actor_milestone_points_redeem_history_records"
