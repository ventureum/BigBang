package actor_profile_record_config

const TABLE_SCHEMA_FOR_ACTOR_PROFILE_RECORD = `
CREATE TABLE actor_profile_records (
    actor TEXT NOT NULL,
    actor_type actor_type_enum NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
CREATE INDEX actor_profile_records_index ON actor_profile_records (actor, actor_type);
`

const TABLE_NAME_FOR_ACTOR_PROFILE_RECORD = "actor_profile_records"
