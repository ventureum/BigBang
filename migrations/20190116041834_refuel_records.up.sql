CREATE TABLE refuel_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    fuel BIGINT NOT NULL DEFAULT 0,
    reputation BIGINT NOT NULL DEFAULT 0,
    milestone_points BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
CREATE INDEX refuel_records_index ON refuel_records (actor, fuel, reputation, milestone_points);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON refuel_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
