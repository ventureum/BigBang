CREATE TABLE actor_milestone_points_redeem_history_records (
    id TEXT NOT NULL,
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
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
    PRIMARY KEY (id)
);
CREATE INDEX actor_milestone_points_redeem_history_records_index ON actor_milestone_points_redeem_history_records (
     actor, redeem_block, redeemed_tokens, submitted_at, executed_at);
CREATE INDEX actor_milestone_points_redeem_history_records_id_desc_index ON actor_milestone_points_redeem_history_records (id DESC NULLS LAST);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON actor_milestone_points_redeem_history_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
