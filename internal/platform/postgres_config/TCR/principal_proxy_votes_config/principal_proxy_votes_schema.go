package principal_proxy_votes_config

const TABLE_SCHEMA_FOR_PRINCIPAL_PROXY_VOTES = `
CREATE TABLE principal_proxy_votes (
    id TEXT NOT NULL,
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    project_id TEXT NOT NULL,
    proxy TEXT NOT NULL,
    block_timestamp BIGINT NOT NULL,
    votes_in_percent BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT principal_proxy_votes_PK
        PRIMARY KEY (actor, project_id, proxy),

    CONSTRAINT actor_actor_delegate_votes_accounts_principal_proxy_votes_FK
       FOREIGN KEY (actor, project_id) 
       REFERENCES actor_delegate_votes_accounts (actor, project_id) ON DELETE CASCADE,

    CONSTRAINT proxy_actor_delegate_votes_accounts_principal_proxy_votes_FK
       FOREIGN KEY (proxy, project_id) 
       REFERENCES actor_delegate_votes_accounts (actor, project_id) ON DELETE CASCADE
);
CREATE INDEX principal_proxy_votes_index ON principal_proxy_votes (actor, project_id, proxy);
CREATE INDEX principal_proxy_votes_desc_index ON proxies (id DESC NULLS LAST);
`

const TABLE_NAME_FOR_PRINCIPAL_PROXY_VOTES = "principal_proxy_votes"
