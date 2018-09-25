package actor_profile_record_config

const TABLE_SCHEMA_FOR_ACTOR_PROFILE_RECORD = `
CREATE TABLE actor_profile_records (
    actor TEXT NOT NULL,
    actor_type actor_type_enum NOT NULL,
    username TEXT NOT NULL,
    photo_url TEXT,
    telegram_id TEXT,
    phone_number TEXT,
    actor_profile_status actor_profile_status_enum NOT NULL DEFAULT 'ACTIVATED',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
CREATE INDEX actor_profile_records_index ON actor_profile_records (actor, actor_type, actor_profile_status, username, created_at, updated_at);
`

const TABLE_NAME_FOR_ACTOR_PROFILE_RECORD = "actor_profile_records"
