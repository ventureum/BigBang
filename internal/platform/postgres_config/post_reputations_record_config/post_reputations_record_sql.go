package post_reputations_record_config

const UPSERT_POST_REPUTATIONS_RECORD_COMMAND = `
WITH counters AS (
  SELECT downvote_count, upvote_count, total_vote_count
  FROM post_reputations_records
  WHERE post_hash = :post_hash AND actor = :actor
)
INSERT INTO post_reputations_records
(
  post_hash,
  actor,
  reputations,
  latest_vote_type,
  downvote_count,
  upvote_count,
  total_vote_count
)
VALUES (
  :post_hash,
  :actor,
  :reputations,
  :latest_vote_type,
  COALESCE((select downvote_count from counters), 0) + CAST((:latest_vote_type = 'DOWN') as integer),
  COALESCE((select upvote_count from counters), 0) + CAST((:latest_vote_type = 'UP') as integer),
  COALESCE((select total_vote_count from counters), 0) + 1
)
ON CONFLICT (post_hash, actor)
DO
  UPDATE
    SET
      reputations = EXCLUDED.reputations,
      latest_vote_type = EXCLUDED.latest_vote_type,
      downvote_count = EXCLUDED.downvote_count,
      upvote_count = EXCLUDED.upvote_count,
      total_vote_count = EXCLUDED.total_vote_count
RETURNING *;
`

const DELETE_POST_REPUTATIONS_RECORDS_BY_POST_HASH_AND_ACTOR_COMMAND = `
DELETE FROM post_reputations_records
WHERE post_hash = $1 and actor = $2;
`

const DELETE_POST_REPUTATIONS_RECORDS_BY_POST_HASH_COMMAND = `
DELETE FROM post_reputations_records
WHERE post_hash = $1;
`

const DELETE_POST_REPUTATIONS_RECORDS_BY_ACTOR_COMMAND = `
DELETE FROM post_reputations_records
WHERE actor = $1;
`

const QUERY_POST_REPUTATIONS_RECORDS_BY_POST_HASH_AND_ACTOR_COMMAND = `
SELECT * FROM post_reputations_records
WHERE post_hash = $1 and actor = $2;
`

const QUERY_TOTAL_REPUTATIONS_BY_POST_HASH_COMMAND = `
SELECT sum(reputations) FROM post_reputations_records
WHERE post_hash = $1;
`

const QUERY_REPUTATIONS_BY_POST_HASH_AND_ACTOR_COMMAND = `
SELECT reputations FROM post_reputations_records
WHERE post_hash = $1 and actor = $2;
`

const QUERY_REPUTATIONS_BY_POST_HASH_AND_VOTE_TYPE_COMMAND = `
SELECT sum(reputations) FROM post_reputations_records
WHERE post_hash = $1 and latest_vote_type = $2;
`

const QUERY_REPUTATIONS_BY_POST_HASH_AND_ACTOR_WITH_LATEST_VOTE_TYPE_AND_TIME_CUTOFF_COMMAND = `
SELECT reputations FROM post_reputations_records
WHERE post_hash = $1 and actor = $2 and latest_vote_type = $3 and updated_at < $4;
`

const QUERY_TOTAL_VOTES_COUNT_BY_POST_HASH_AND_ACTOR_COMMAND = `
SELECT total_vote_count FROM post_reputations_records
WHERE post_hash = $1 and actor = $2;
`

const QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND = `
SELECT actor FROM post_reputations_records
WHERE post_hash = $1 and latest_vote_type = $2;
`