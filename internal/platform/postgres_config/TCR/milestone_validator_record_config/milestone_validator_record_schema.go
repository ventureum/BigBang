package milestone_validator_record_config

const TABLE_SCHEMA_FOR_MILESTONE_VALIDATOR_RECORD = `
CREATE TABLE milestone_validator_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    project_id TEXT NOT NULL,
    milestone_id BIGINT NOT NULL,
    validator TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT milestone_validator_records_PK
        PRIMARY KEY (uuid),

    CONSTRAINT milestone_milestone_validator_record_FK
       FOREIGN KEY (project_id, milestone_id) 
       REFERENCES milestones (project_id, milestone_id) ON DELETE CASCADE,

    CONSTRAINT validator_actor_profile_records_FK
       FOREIGN KEY (validator) 
       REFERENCES actor_profile_records (actor) ON DELETE CASCADE
);
CREATE INDEX milestone_validator_records_index ON milestone_validator_records (project_id, milestone_id, validator);
`

const TABLE_NAME_FOR_MILESTONE_VALIDATOR_RECORD = "milestone_validator_records"
