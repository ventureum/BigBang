package rating_vote_config


const TABLE_SCHEMA_FOR_RATING_VOTE = `
CREATE TABLE rating_votes (
    project_id TEXT NOT NULL,
    milestone_id BIGINT NOT NULL,
    objective_id BIGINT NOT NULL,
    voter TEXT NOT NULL,
    rating BIGINT NOT NULL DEFAULT 0,
    weight BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (project_id, milestone_id, objective_id, voter)
);
CREATE INDEX rating_votes_index ON rating_votes(project_id, milestone_id, objective_id, voter);
CREATE INDEX rating_votes_asc_index ON rating_votes (created_at ASC NULLS LAST);
`

const TABLE_NAME_FOR_RATING_VOTE = "rating_votes"
