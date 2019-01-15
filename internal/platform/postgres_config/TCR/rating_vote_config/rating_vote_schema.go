package rating_vote_config

const TABLE_SCHEMA_FOR_RATING_VOTE = `
CREATE TABLE rating_votes (
    id TEXT NOT NULL,
    project_id TEXT NOT NULL,
    milestone_id BIGINT NOT NULL,
    objective_id BIGINT NOT NULL,
    voter TEXT NOT NULL,
    block_timestamp BIGINT NOT NULL,
    rating BIGINT NOT NULL DEFAULT 0,
    weight BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT rating_votes_PK
      PRIMARY KEY (project_id, milestone_id, objective_id, voter),

    CONSTRAINT actor_profile_records_rating_votes_FK
       FOREIGN KEY (voter) 
       REFERENCES actor_profile_records (actor),
  
    CONSTRAINT objectives_rating_votes_FK
       FOREIGN KEY (project_id, milestone_id, objective_id) 
       REFERENCES objectives (project_id, milestone_id, objective_id) ON DELETE CASCADE
);
CREATE INDEX rating_votes_index ON rating_votes(id, project_id, milestone_id, objective_id, voter);
CREATE INDEX rating_votes_asc_index ON rating_votes (id DESC NULLS LAST);
`

const TABLE_NAME_FOR_RATING_VOTE = "rating_votes"
