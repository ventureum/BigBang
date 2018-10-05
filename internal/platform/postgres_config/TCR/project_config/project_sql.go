package project_config

const UPSERT_PROJECT_COMMAND = `
INSERT INTO projects 
(
  project_id,
  admin,
  content,
  avg_rating,
  current_milestone,
  num_milestones,
  num_milestones_completed
)
VALUES 
(
  :project_id,
  :admin,
  :content,
  :avg_rating,
  :current_milestone,
  :num_milestones,
  :num_milestones_completed
)
ON CONFLICT (project_id) 
DO
 UPDATE
    SET
        admin = :admin,
        content = :content,
        avg_rating = :avg_rating,
        current_milestone = :current_milestone,
        num_milestones = :num_milestones,
        num_milestones_completed = :num_milestones_completed
    WHERE projects.project_id = :project_id
RETURNING created_at;
`

const DELETE_PROJECT_COMMAND = `
DELETE FROM projects
WHERE project_id = $1;
`

const QUERY_PROJECT_COMMAND = `
SELECT * FROM projects
WHERE project_id = $1;
`

const VERIFY_PROJECT_EXISTING_COMMAND = `
select exists(select 1 from projects where project_id =$1);
`

const PAGINATION_QUERY_PROJECT_LIST_COMMAND = `
SELECT * FROM projects
WHERE id <= $1
ORDER BY id DESC
LIMIT $2;
`

const QUERY_PROJECT_LIST_COMMAND = `
SELECT * FROM projects
ORDER BY id DESC
LIMIT $1;
`
