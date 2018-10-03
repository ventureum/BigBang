package project_config

const UPSERT_PROJECT_COMMAND = `
INSERT INTO projects 
(
  project_id,
  content,
  avg_rating,
  milestone_info
)
VALUES 
(
  :project_id,
  :content,
  :avg_rating,
  :milestone_info
)
ON CONFLICT (project_id) 
DO
 UPDATE
    SET
        content = :content,
        avg_rating = :avg_rating,
        milestone_info = :milestone_info
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
