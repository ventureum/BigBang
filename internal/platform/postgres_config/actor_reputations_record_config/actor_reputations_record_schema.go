package actor_reputations_record_config

const TABLE_SCHEMA_FOR_ACTOR_REPUTATIONS_RECORD = `
CREATE TABLE actor_reputations_records (
    actor TEXT NOT NULL,
    reputations BIGINT NOT NULL DEFAULT 0 check(reputations >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
CREATE INDEX actor_reputations_records_index ON actor_reputations_records (actor, reputations);
`

const TABLE_NAME_FOR_ACTOR_REPUTATIONS_RECORD = "actor_reputations_records"
