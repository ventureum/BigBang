package reputations_refuel_record_config

const TABLE_SCHEMA_FOR_REPUTATION_REFUEL_RECORD = `
CREATE TABLE reputations_refuel_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    actor TEXT NOT NULL,
    reputations BIGINT NOT NULL DEFAULT 0 check(reputations >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
CREATE INDEX reputations_refuel_records_index ON reputations_refuel_records (actor, reputations);
`

const TABLE_NAME_FOR_REPUTATION_REFUEL_RECORD = "reputations_refuel_records"
