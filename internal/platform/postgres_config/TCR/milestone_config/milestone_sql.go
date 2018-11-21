package milestone_config

const UPSERT_MILESTONE_COMMAND = `
INSERT INTO milestones 
(
  project_id,
  milestone_id,
  content,
  block_timestamp,
  start_time,
  end_time,
  state
)
VALUES 
(
  :project_id,
  :milestone_id,
  :content,
  :block_timestamp,
  :start_time,
  :end_time,
  :state
)
ON CONFLICT (project_id, milestone_id) 
DO
 UPDATE
    SET
        content = :content,
        start_time = :start_time,
        end_time = :end_time,
        state = :state
    WHERE milestones.project_id = :project_id and milestones.milestone_id = :milestone_id
RETURNING (xmax = 0) AS inserted;
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

const INVALID_MILESTONE_UPDATE_IF_EXISTING_COMMAND = `
select exists(select 1 from milestones where project_id = $1 and milestone_id = $2 and state != 'PENDING');
`

const INCREASE_NUM_OBJECTIVES_COMMAND = `
UPDATE milestones
SET
   num_objectives = num_objectives + 1
WHERE project_id = $1 and milestone_id = $2;
`

const DECREASE_NUM_OBJECTIVES_COMMAND = `
UPDATE milestones
SET
   num_objectives = num_objectives - 1
WHERE project_id = $1 and milestone_id = $2;
`

const ADD_RATING_AND_WEIGHT_FOR_MILESTONE_COMMAND = `
UPDATE milestones
SET
   total_rating = total_rating + $3,
   total_weight = total_weight + $4,
   avg_rating = (total_rating + $3) / GREATEST(total_weight + $4, 1)
WHERE project_id = $1 and milestone_id = $2;
`

const ACTIVATE_MILESTONE_COMMAND = `
UPDATE milestones
SET
  block_timestamp = $3,
  start_time = $4,
  state = 'IN_PROGRESS'
WHERE project_id = $1 and milestone_id = $2;
`

const FINALIZE_MILESTONE_COMMAND = `
UPDATE milestones
SET
  block_timestamp = $3,
  end_time = $4,
  state = 'COMPLETE'
WHERE project_id = $1 and milestone_id = $2;
`
