package actor_delegate_votes_account_config

const TABLE_SCHEMA_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT = `
CREATE TABLE actor_delegate_votes_accounts (
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    project_id TEXT NOT NULL,
    available_delegate_votes BIGINT NOT NULL DEFAULT 0,
    received_delegate_votes BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT actor_delegate_votes_accounts_PK
        PRIMARY KEY (actor, project_id),

    CONSTRAINT actor_profile_records_actor_delegate_votes_accounts_FK
       FOREIGN KEY (actor) 
       REFERENCES actor_profile_records (actor),

    CONSTRAINT projects_actor_delegate_votes_accounts_FK
       FOREIGN KEY (project_id) 
       REFERENCES projects (project_id) ON DELETE CASCADE
);
CREATE INDEX actor_delegate_votes_accounts_index ON actor_delegate_votes_accounts (actor, project_id);
`

const TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT = "actor_delegate_votes_accounts"
