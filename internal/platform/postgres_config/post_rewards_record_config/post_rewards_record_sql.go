package post_rewards_record_config


const UPSERT_POST_REWARDS_RECORD_COMMAND = `
INSERT INTO post_rewards_records
(
  post_hash,
  rewards
)
VALUES 
(
  :post_hash, 
  :rewards
)
ON CONFLICT (post_hash) 
DO
 UPDATE
    SET rewards = :rewards
    WHERE post_rewards_records.post_hash = :post_hash;
`

const DELETE_POST_REWARDS_RECORD_COMMAND = `
DELETE FROM post_rewards_records
WHERE post_hash = $1;
`

const QUERY_POST_REWARDS_COMMAND = `
SELECT rewards FROM post_rewards_records
WHERE post_hash = $1;
`

const ADD_POST_REWARDS_COMMAND = `
 UPDATE post_rewards_records
    SET rewards = post_rewards_records.rewards + $2
    WHERE post_hash = $1;
`

const SUB_POST_REWARDS_COMMAND = `
 UPDATE post_rewards_records
    SET rewards = post_rewards_records.rewards - $2
    WHERE post_hash = $1;
`

const UPSERT_POST_REWARDS_RECORD_BY_AGGREGATION_COMMAND = `
with updates as (
    SELECT
      post_hash,
      0.001 * count(*) * percentile_cont(0.5) within group (order by signed_reputations) as median_reputations
    FROM
      post_votes_records
    GROUP BY
      post_hash
)

INSERT INTO post_rewards_records
SELECT post_hash,  GREATEST(median_reputations, 0), median_reputations From updates
ON CONFLICT (post_hash)
  DO
  UPDATE
       SET rewards = GREATEST(EXCLUDED.rewards,  post_rewards_records.rewards),
           latest_median_reputations = EXCLUDED.latest_median_reputations,
           withdrawable_rewards = post_rewards_records.withdrawable_rewards + GREATEST(EXCLUDED.rewards - post_rewards_records.rewards, 0);
`
