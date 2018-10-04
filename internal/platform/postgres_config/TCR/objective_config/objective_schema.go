package objective_config


const TABLE_SCHEMA_FOR_OBJECTIVE = `
CREATE TABLE objectives (
    project_id TEXT NOT NULL,
    milestone_id BIGINT NOT NULL,
    objective_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    total_rating BIGINT NOT NULL DEFAULT 0,
    total_weight BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (project_id, milestone_id, objective_id)
);
CREATE INDEX objectives_index ON objectives (project_id, milestone_id, objective_id);
CREATE INDEX objectives_asc_index ON objectives (objective_id ASC NULLS LAST);
`

const TABLE_NAME_FOR_OBJECTIVE = "objectives"
