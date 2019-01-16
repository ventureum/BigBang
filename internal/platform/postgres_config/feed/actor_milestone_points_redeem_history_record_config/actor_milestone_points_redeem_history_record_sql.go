package actor_milestone_points_redeem_history_record_config

const UPSERT_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD_COMMAND = `
INSERT INTO actor_milestone_points_redeem_history_records
(   
    id,
    actor,
    redeem_block,
    token_pool,
    total_enrolled_milestone_points,
    targeted_milestone_points,
    actual_milestone_points,
    consumed_milestone_points,
    redeemed_tokens,
    submitted_at,
    executed_at
)
VALUES 
(
    :id,
    :actor,
    :redeem_block,
    :token_pool,
    :total_enrolled_milestone_points,
    :targeted_milestone_points,
    :actual_milestone_points,
    :consumed_milestone_points,
    :redeemed_tokens,
    :submitted_at,
    :executed_at
);
`

const DELETE_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORD_BY_ACTOR_COMMAND = `
DELETE FROM actor_milestone_points_redeem_history_records
WHERE actor = $1;
`

const QUERY_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORDS_COMMAND = `
SELECT actor, redeem_block, tokenPool, total_enrolled_milestone_points,  targeted_milestone_points, actual_milestone_points, 
consumed_milestone_points, redeemed_tokens, submitted_at, executed_at
FROM actor_milestone_points_redeem_history_records
WHERE actor = $1;
`

const UPSERT_BATCH_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_RECORDS_BY_REDEEM_BLOCK = `
with updates as(
  INSERT INTO actor_milestone_points_redeem_history_records(
    SELECT
      concat(milestone_points_redeem_request_records.next_redeem_block::text, ':', milestone_points_redeem_request_records.actor) as id,
      milestone_points_redeem_request_records.actor as actor,
      milestone_points_redeem_request_records.next_redeem_block as redeem_block,
      redeem_block_info_records.token_pool as token_pool,
      redeem_block_info_records.total_enrolled_milestone_points as total_enrolled_milestone_points,
      milestone_points_redeem_request_records.targeted_milestone_points as targeted_milestone_points,
      actor_rewards_info_records.milestone_points as actual_milestone_points,
      LEAST(milestone_points_redeem_request_records.targeted_milestone_points,  actor_rewards_info_records.milestone_points) as consumed_milestone_points,
      (redeem_block_info_records.token_pool *  LEAST(milestone_points_redeem_request_records.targeted_milestone_points,  actor_rewards_info_records.milestone_points) / GREATEST(redeem_block_info_records.total_enrolled_milestone_points, 1)) as redeemed_tokens,
      milestone_points_redeem_request_records.updated_at as submitted_at,
      redeem_block_info_records.executed_at as executed_at
    FROM milestone_points_redeem_request_records, actor_rewards_info_records, redeem_block_info_records
    WHERE milestone_points_redeem_request_records.actor = actor_rewards_info_records.actor and redeem_block_info_records.redeem_block = $1 and
          (milestone_points_redeem_request_records.next_redeem_block = $1 or milestone_points_redeem_request_records.targeted_milestone_points = 9223372036854775807)
  ) RETURNING actor as actor, consumed_milestone_points as consumed_milestone_points)
UPDATE  actor_rewards_info_records
SET
  milestone_points = actor_rewards_info_records.milestone_points - updates.consumed_milestone_points,
  consumed_milestone_points = actor_rewards_info_records.consumed_milestone_points + updates.consumed_milestone_points
FROM updates
WHERE actor_rewards_info_records.actor = updates.actor;
`

const VERIFY_REDEEM_BLOCK_EXISTING_COMMAND = `
select exists(select 1 from actor_milestone_points_redeem_history_records where redeem_block = $1);
`

const VERIFY_REDEEM_BLOCK_FOR_ACTOR_EXISTING_COMMAND = `
select exists(select 1 from actor_milestone_points_redeem_history_records where actor = $1 and redeem_block = $2);
`

const PAGINATION_QUERY_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_COMMAND = `
SELECT actor, redeem_block, token_pool, total_enrolled_milestone_points,  targeted_milestone_points, actual_milestone_points, 
consumed_milestone_points, redeemed_tokens, submitted_at, executed_at
FROM actor_milestone_points_redeem_history_records
WHERE actor = $1 and id <= $2 
ORDER BY id DESC
LIMIT $3
`

const QUERY_ACTOR_MILESTONE_POINTS_REDEEM_HISTORY_WITH_LIMIT_COMMAND = `
SELECT actor, redeem_block, token_pool, total_enrolled_milestone_points, targeted_milestone_points, actual_milestone_points, 
consumed_milestone_points, redeemed_tokens, submitted_at, executed_at
FROM actor_milestone_points_redeem_history_records
WHERE actor = $1
ORDER BY id DESC
LIMIT $2;
`
