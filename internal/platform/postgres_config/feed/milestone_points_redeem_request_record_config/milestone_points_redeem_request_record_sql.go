package milestone_points_redeem_request_record_config

const UPSERT_MILESTONE_POINTS_REDEEM_REQUEST_RECORD_COMMAND = `
INSERT INTO milestone_points_redeem_request_records
(
  actor,
  next_redeem_block,
  targeted_milestone_points
)
VALUES 
(
  :actor,
  :next_redeem_block,
  :targeted_milestone_points
)
ON CONFLICT (actor) 
DO
 UPDATE
    SET 
        next_redeem_block = :next_redeem_block,
        targeted_milestone_points = :targeted_milestone_points
    WHERE milestone_points_redeem_request_records.actor = :actor;
`

const DELETE_MILESTONE_POINTS_REDEEM_REQUEST_RECORD_COMMAND = `
DELETE FROM milestone_points_redeem_request_records
WHERE actor = $1;
`

const QUERY_MILESTONE_POINTS_REDEEM_REQUEST_COMMAND = `
SELECT actor, next_redeem_block, targeted_milestone_points, updated_at FROM milestone_points_redeem_request_records
WHERE actor = $1;
`

const QUERY_TOTAL_ENROLLED_MILESTONE_POINTS_COMMAND = `
SELECT sum(targeted_milestone_points) FROM milestone_points_redeem_request_records
WHERE next_redeem_block = $1 or targeted_milestone_points = 9223372036854775807;
`

const VERIFY_ACTOR_EXISTING_COMMAND = `
select exists(select 1 from milestone_points_redeem_request_records where actor =$1);
`
