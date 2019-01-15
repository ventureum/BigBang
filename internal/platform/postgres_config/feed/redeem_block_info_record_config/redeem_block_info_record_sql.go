package redeem_block_info_record_config

const UPSERT_REDEEM_BLOCK_INFO_RECORD_COMMAND = `
INSERT INTO redeem_block_info_records
(
  redeem_block,
  total_enrolled_milestone_points,
  token_pool
)
VALUES 
(
  :redeem_block,
  :total_enrolled_milestone_points,
  :token_pool
)
ON CONFLICT (redeem_block) 
DO
 UPDATE
    SET 
        total_enrolled_milestone_points = :total_enrolled_milestone_points,
        token_pool = :token_pool
    WHERE redeem_block_info_records.redeem_block = :redeem_block;
`

const INIT_REDEEM_BLOCK_INFO_RECORD_COMMAND = `
INSERT INTO redeem_block_info_records
(
  redeem_block,
  executed_at 
)
VALUES 
(
  $1,
  $2
)
ON CONFLICT (redeem_block) 
DO NOTHING;
`

const DELETE_REDEEM_BLOCK_INFO_RECORD_COMMAND = `
DELETE FROM redeem_block_info_records
WHERE redeem_block = $1 and executed_at = $2;
`

const QUERY_REDEEM_BLOCK_INFO_COMMAND = `
SELECT redeem_block, total_enrolled_milestone_points, token_pool, executed_at FROM redeem_block_info_records
WHERE redeem_block = $1;
`

const UPDATE_EXECUTED_AT_COMMAND = `
 UPDATE redeem_block_info_records
    SET 
        executed_at = NOW()
    WHERE redeem_block_info_records.redeem_block = $1;
`

const UPDATE_TOOKEN_POOL_COMMAND = `
 UPDATE redeem_block_info_records
    SET 
        token_pool = $2
    WHERE redeem_block = $1;
`

const UPDATE_TOTAL_ENROLLED_MILESTONEPOINTS_COMMAND = `
with updates as (
    SELECT  
         COALESCE(
           sum(LEAST(milestone_points_redeem_request_records.targeted_milestone_points, actor_rewards_info_records.milestone_points)), 0) as total_enrolled_milestone_points 
    FROM milestone_points_redeem_request_records, actor_rewards_info_records
    WHERE milestone_points_redeem_request_records.actor = actor_rewards_info_records.actor and
          (milestone_points_redeem_request_records.next_redeem_block = $1 or milestone_points_redeem_request_records.targeted_milestone_points = 9223372036854775807))
UPDATE redeem_block_info_records
  SET total_enrolled_milestone_points = updates.total_enrolled_milestone_points
FROM updates
WHERE  redeem_block_info_records.redeem_block = $1;
`

const VERIFY_REDEEM_BLOCK_INFO_RECORD_EXISTING_COMMAND = `
select exists(select 1 from redeem_block_info_records where redeem_block = $1);
`
