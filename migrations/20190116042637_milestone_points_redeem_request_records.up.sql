CREATE TABLE milestone_points_redeem_request_records (
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    next_redeem_block BIGINT NOT NULL DEFAULT 0,
    targeted_milestone_points BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
CREATE INDEX milestone_points_redeem_request_records_index ON milestone_points_redeem_request_records (actor, next_redeem_block, targeted_milestone_points);
CREATE INDEX total_enrolled_milestone_points_index ON milestone_points_redeem_request_records (next_redeem_block, targeted_milestone_points);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON milestone_points_redeem_request_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
