package project_config

const TABLE_SCHEMA_FOR_PROJECT = `
CREATE TABLE projects (
    id TEXT NOT NULL UNIQUE,
    project_id TEXT NOT NULL,
    admin Text NOT NULL,
    content TEXT NOT NULL,
    block_timestamp BIGINT NOT NULL,
    avg_rating BIGINT NOT NULL DEFAULT 0,
    total_rating BIGINT NOT NULL DEFAULT 0,
    total_weight BIGINT NOT NULL DEFAULT 0,
    current_milestone BIGINT NOT NULL DEFAULT 0,
    num_milestones BIGINT NOT NULL DEFAULT 0,
    num_milestones_completed BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (project_id)
);
CREATE INDEX projects_index ON projects (project_id);
CREATE INDEX projects_desc_index ON projects (id DESC);
`

const TABLE_NAME_FOR_PROJECT = "projects"
