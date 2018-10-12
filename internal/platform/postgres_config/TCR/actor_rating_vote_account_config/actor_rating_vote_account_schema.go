package actor_rating_vote_account_config


const TABLE_SCHEMA_FOR_ACTOR_RATING_VOTE_ACCOUNT = `
CREATE TABLE actor_rating_vote_accounts (
    actor TEXT NOT NULL,
    project_id TEXT NOT NULL,
    available_rating_votes BIGINT NOT NULL DEFAULT 0,
    received_rating_votes BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT actor_rating_vote_accounts_PK
        PRIMARY KEY (actor, project_id),

    CONSTRAINT actor_profile_records_actor_rating_vote_accounts_FK
       FOREIGN KEY (actor) 
       REFERENCES actor_profile_records (actor),

    CONSTRAINT projects_actor_rating_vote_accounts_FK
       FOREIGN KEY (project_id) 
       REFERENCES projects (project_id) ON DELETE CASCADE
);
CREATE INDEX actor_rating_vote_accounts_index ON actor_rating_vote_accounts (actor, project_id);
`

const TABLE_NAME_FOR_ACTOR_RATING_VOTE_ACCOUNT = "actor_rating_vote_accounts"
