package project_config

const INSERT_PROJECT_COMMAND = `
INSERT INTO projects 
(
  project_id,
  admin,
  content
)
VALUES 
(
  :project_id,
  :admin,
  :content
);
`

const UPDATE_PROJECT_COMMAND = `
UPDATE projects
    SET
        admin = :admin,
        content = :content
WHERE projects.project_id = :project_id;
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

const ADD_RATING_AND_WEIGHT_COMMAND = `
UPDATE projects
SET 
   total_rating = total_rating + $2,
   total_weight = total_weight + $3,
   avg_rating = (total_rating + $2) / GREATEST(total_weight + $3, 1)
WHERE project_id = $1
`

const INCREASE_NUM_MILESTONES_COMMAND = `
UPDATE projects
SET
   num_milestones = num_milestones + 1
WHERE project_id = $1
`

const INCREASE_NUM_MILESTONES_COMPLETED_COMMAND = `
UPDATE projects
SET
   num_milestones_completed = LEAST(num_milestones_completed + 1, num_milestones),
   current_milestone = 0
WHERE project_id = $1
`

const SET_CURRENT_MILESTONE_COMMAND = `
UPDATE projects
SET
   current_milestone = (num_milestones_completed + 1)%(num_milestones + 1)
WHERE project_id = $1
`
