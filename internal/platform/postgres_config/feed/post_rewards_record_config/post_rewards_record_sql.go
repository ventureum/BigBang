package post_rewards_record_config

const UPSERT_POST_REWARDS_RECORD_COMMAND = `
INSERT INTO post_rewards_records
(
  post_hash,
  actor,
  post_type,
  object,
  post_time,
  delta_fuel
)
VALUES 
(
  :post_hash,
  :actor,
  :post_type,
  :object,
  :post_time,
  :delta_fuel
)
ON CONFLICT (post_hash) 
DO
 UPDATE
    SET  
       actor = :actor,
       post_type = :post_type,
       delta_fuel = post_rewards_records.delta_fuel + :delta_fuel,
       object = :object,
       post_time = :post_time
    WHERE post_rewards_records.post_hash = :post_hash;
`

const DELETE_POST_REWARDS_RECORD_COMMAND = `
DELETE FROM post_rewards_records
WHERE post_hash = $1;
`

const QUERY_POST_REWARDS_RECORD_COMMAND = `
SELECT * FROM post_rewards_records
WHERE post_hash = $1;
`

const UPSERT_POST_REWARDS_RECORD_BY_AGGREGATION_COMMAND = `
with updates as (
    SELECT
      post_hash,
      0.01 * count(*) * percentile_cont(0.5) within group (order by signed_reputation) as delta_reputation
    FROM
      post_votes_records
    GROUP BY
      post_hash)

UPDATE post_rewards_records
  SET withdrawable_mps = post_rewards_records.withdrawable_mps +  updates.delta_reputation - post_rewards_records.delta_milestone_points,
    delta_reputation =  updates.delta_reputation,
    delta_milestone_points =  updates.delta_reputation
FROM updates
WHERE  post_rewards_records.post_hash = updates.post_hash
RETURNING post_rewards_records.object, post_rewards_records.post_time, post_rewards_records.withdrawable_mps;
`

const QUERY_RECENT_POST_REWARDS_RECORDS_BY_ACTOR_COMMAND = `
SELECT * FROM post_rewards_records
WHERE actor = $1 and post_type = $2 ORDER BY created_at DESC LIMIT $3;
`
