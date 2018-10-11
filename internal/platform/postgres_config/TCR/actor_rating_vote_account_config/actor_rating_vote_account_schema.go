package actor_rating_vote_account_config


const TABLE_SCHEMA_FOR_ACTOR_RATING_VOTE_ACCOUNT = `
CREATE TABLE actor_rating_vote_accounts (
    actor TEXT NOT NULL REFERENCES actor_profile_records (actor),
    available_rating_votes BIGINT NOT NULL DEFAULT 0,
    received_rating_votes BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (actor)
);
`

const TABLE_NAME_FOR_ACTOR_RATING_VOTE_ACCOUNT = "actor_rating_vote_accounts"
