package actor_votes_counters_record_config

const UPSERT_ACTOR_VOTES_COUNTERS_RECORD_COMMAND = `
WITH counters AS (
  SELECT latest_reputation_for_upvote, latest_reputation_for_downvote, downvote_count, upvote_count, total_vote_count
  FROM actor_votes_counters_records
  WHERE post_hash = :post_hash AND actor = :actor
)
INSERT INTO actor_votes_counters_records
(
  post_hash,
  actor,
  latest_reputation,
  latest_vote_type,
  latest_reputation_for_upvote,
  latest_reputation_for_downvote,
  downvote_count,
  upvote_count,
  total_vote_count
)
VALUES (
  :post_hash,
  :actor,
  :latest_reputation,
  :latest_vote_type,
  COALESCE((select latest_reputation_for_upvote from counters), 0) * CAST((:latest_vote_type = 'DOWN') as integer) +  :latest_reputation * CAST((:latest_vote_type = 'UP') as integer),
  COALESCE((select latest_reputation_for_downvote from counters), 0) * CAST((:latest_vote_type = 'UP') as integer) +  :latest_reputation * CAST((:latest_vote_type = 'DOWN') as integer),
  COALESCE((select downvote_count from counters), 0) + CAST((:latest_vote_type = 'DOWN') as integer),
  COALESCE((select upvote_count from counters), 0) + CAST((:latest_vote_type = 'UP') as integer),
  COALESCE((select total_vote_count from counters), 0) + 1
)
ON CONFLICT (post_hash, actor)
DO
  UPDATE
    SET
      latest_reputation = EXCLUDED.latest_reputation,
      latest_vote_type = EXCLUDED.latest_vote_type,
      latest_reputation_for_upvote = EXCLUDED.latest_reputation_for_upvote,
      latest_reputation_for_downvote = EXCLUDED.latest_reputation_for_downvote,
      downvote_count = EXCLUDED.downvote_count,
      upvote_count = EXCLUDED.upvote_count,
      total_vote_count = EXCLUDED.total_vote_count
RETURNING *;
`

const DELETE_ACTOR_VOTES_COUNTERS_RECORDS_BY_POST_HASH_AND_ACTOR_COMMAND = `
DELETE FROM actor_votes_counters_records
WHERE post_hash = $1 and actor = $2;
`

const DELETE_ACTOR_VOTES_COUNTERS_RECORDS_BY_POST_HASH_COMMAND = `
DELETE FROM actor_votes_counters_records
WHERE post_hash = $1;
`

const DELETE_ACTOR_VOTES_COUNTERS_RECORDS_BY_ACTOR_COMMAND = `
DELETE FROM actor_votes_counters_records
WHERE actor = $1;
`

const QUERY_ACTOR_VOTES_COUNTERS_RECORD_BY_POST_HASH_AND_ACTOR_COMMAND = `
SELECT * FROM actor_votes_counters_records
WHERE post_hash = $1 and actor = $2;
`

const QUERY_TOTAL_REPUTATION_BY_POST_HASH_COMMAND = `
SELECT sum(latest_reputation) FROM actor_votes_counters_records
WHERE post_hash = $1;
`

const QUERY_REPUTATION_BY_POST_HASH_AND_ACTOR_COMMAND = `
SELECT latest_reputation FROM actor_votes_counters_records
WHERE post_hash = $1 and actor = $2;
`

const QUERY_REPUTATION_BY_POST_HASH_AND_VOTE_TYPE_COMMAND = `
SELECT sum(latest_reputation) FROM actor_votes_counters_records
WHERE post_hash = $1 and latest_vote_type = $2;
`

const QUERY_REPUTATION_BY_POST_HASH_AND_ACTOR_WITH_LATEST_VOTE_TYPE_AND_TIME_CUTOFF_COMMAND = `
SELECT latest_reputation_for_upvote * CAST((latest_vote_type = 'UP') as integer) +  latest_reputation_for_downvote * CAST((latest_vote_type = 'DOWN') as integer) FROM actor_votes_counters_records
WHERE post_hash = $1 and actor = $2 and latest_vote_type = $3 and updated_at < $4;
`

const QUERY_TOTAL_VOTES_COUNT_BY_POST_HASH_AND_ACTOR_COMMAND = `
SELECT total_vote_count FROM actor_votes_counters_records
WHERE post_hash = $1 and actor = $2;
`

const QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND = `
SELECT actor FROM actor_votes_counters_records
WHERE post_hash = $1 and latest_vote_type = $2;
`

const QUARY_TOTAL_REPUTATION_BY_POSTHASH_COMMAND = `
SELECT COALESCE(sum(latest_reputation), 0) FROM actor_votes_counters_records
WHERE post_hash = $1;
`
