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
    WHERE redeem_block_info_records.actor = :actor;
`

const DELETE_REDEEM_BLOCK_INFO_RECORD_COMMAND = `
DELETE FROM redeem_block_info_records
WHERE redeem_block = $1;
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

const UPDATE_TOTAL_ENROLLED_MILESTONEPOINTS_COMMAND = `
with updates as (
  SELECT next_redeem_block, COALESCE(sum(targeted_milestone_points), 0) as total_enrolled_milestone_points FROM milestone_points_redeem_request_records
  WHERE next_redeem_block = $1 or targeted_milestone_points = 9223372036854775807)

UPDATE redeem_block_info_records
  SET total_enrolled_milestone_points = updates.total_enrolled_milestone_points
FROM updates
WHERE  redeem_block_info_records.redeem_block = updates.next_redeem_block
`

const VERIFY_REDEEM_BLOCK_INFO_RECORD_EXISTING_COMMAND = `
select exists(select 1 from redeem_block_info_records where redeem_block =$1);
`
