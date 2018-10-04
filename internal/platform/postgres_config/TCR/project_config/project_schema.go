package project_config


const TABLE_SCHEMA_FOR_PROJECT = `
CREATE TABLE projects (
    id BIGSERIAL,
    project_id TEXT NOT NULL,
    content TEXT NOT NULL,
    avg_rating BIGINT NOT NULL DEFAULT 0,
    milestone_info JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (project_id)
);
CREATE INDEX projects_index ON projects (id, project_id, avg_rating);
CREATE INDEX projects_desc_index ON projects (id DESC NULLS LAST);
`

const TABLE_NAME_FOR_PROJECT = "projects"
