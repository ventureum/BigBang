package milestone_config


const TABLE_SCHEMA_FOR_MILESTONE = `
CREATE TABLE milestones (
    project_id TEXT NOT NULL,
    milestone_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    start_time BIGINT NOT NULL,
    end_time BIGINT NOT NULL,
    num_objs BIGINT NOT NULL DEFAULT 0,
    avg_rating BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (project_id, milestone_id)
);
CREATE INDEX milestones_index ON milestones (project_id, milestone_id);
CREATE INDEX milestones_asc_index ON milestones (milestone_id ASC NULLS LAST);
`

const TABLE_NAME_FOR_MILESTONE = "milestones"
