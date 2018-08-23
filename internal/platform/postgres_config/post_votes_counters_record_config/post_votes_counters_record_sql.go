package post_votes_counters_record_config

const UPSERT_POST_VOTES_COUNTRS_RECORD_COMMAND = `
WITH counters AS (
  SELECT downvote_count, upvote_count, total_vote_count
  FROM post_votes_counters_records
  WHERE post_hash = :post_hash
)
INSERT INTO post_votes_counters_records
(
  post_hash,
  latest_vote_type,
  downvote_count,
  upvote_count,
  total_vote_count
)
VALUES (
  :post_hash,
  :latest_vote_type,
  COALESCE((select downvote_count from counters), 0) + CAST((:latest_vote_type = 'DOWN') as integer),
  COALESCE((select upvote_count from counters), 0) + CAST((:latest_vote_type = 'UP') as integer),
  COALESCE((select total_vote_count from counters), 0) + 1
)
ON CONFLICT (post_hash)
DO
  UPDATE
    SET
      latest_vote_type = EXCLUDED.latest_vote_type,
      downvote_count = EXCLUDED.downvote_count,
      upvote_count = EXCLUDED.upvote_count,
      total_vote_count = EXCLUDED.total_vote_count
RETURNING *;
`

const DELETE_POST_VOTES_COUNTRS_RECORDS_BY_POST_HASH_COMMAND = `
DELETE FROM post_votes_counters_records
WHERE post_hash = $1;
`
const QUERY_POST_VOTES_COUNTRS_RECORDS_BY_POST_HASH_COMMAND = `
SELECT * FROM post_votes_counters_records
WHERE post_hash = $1;
`
