package milestone_points_redeem_request_record_config

const TABLE_SCHEMA_FOR_MILESTONE_POINTS_REDEEM_REQUEST_RECORD = `
CREATE TABLE milestone_points_redeem_request_records (
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor),
    next_redeem_block BIGINT NOT NULL DEFAULT 0,
    targeted_milestone_points BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
CREATE INDEX milestone_points_redeem_request_records_index ON milestone_points_redeem_request_records (actor, next_redeem_block, targeted_milestone_points);
CREATE INDEX total_enrolled_milestone_points_index ON milestone_points_redeem_request_records (next_redeem_block, targeted_milestone_points);
`

const TABLE_NAME_FOR_MILESTONE_POINTS_REDEEM_REQUEST_RECORD = "milestone_points_redeem_request_records"
