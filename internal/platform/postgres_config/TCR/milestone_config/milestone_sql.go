package milestone_config

const UPSERT_MILESTONE_COMMAND = `
INSERT INTO milestones 
(
  project_id,
  milestone_id,
  content,
  start_time,
  end_time,
  num_objs,
  avg_rating
)
VALUES 
(
  :project_id,
  :milestone_id,
  :content,
  :start_time,
  :end_time,
  :num_objs,
  :avg_rating
)
ON CONFLICT (project_id, milestone_id) 
DO
 UPDATE
    SET
        content = :content,
        start_time = :start_time,
        end_time = :end_time,
        num_objs = :num_objs,
        avg_rating = :avg_rating
    WHERE milestones.project_id = :project_id and milestones.milestone_id = :milestone_id
RETURNING created_at;
`

const DELETE_MILESTONE_BY_IDS_COMMAND = `
DELETE FROM milestones
WHERE project_id = $1 and milestone_id = $2;
`

const DELETE_MILESTONES_BY_PROJECT_ID_COMMAND = `
DELETE FROM milestones
WHERE project_id = $1;
`

const QUERY_MILESTONE_BY_IDS_COMMAND = `
SELECT * FROM milestones
WHERE project_id = $1 and milestone_id = $2;
`

const QUERY_MILESTONES_BY_PROJECT_ID_COMMAND = `
SELECT * FROM milestones
WHERE project_id = $1
ORDER BY milestone_id ASC;
`

const VERIFY_MILESTONE_EXISTING_COMMAND = `
select exists(select 1 from milestones where project_id = $1 and milestone_id = $2);
`
