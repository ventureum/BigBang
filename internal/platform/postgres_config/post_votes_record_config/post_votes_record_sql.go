package post_votes_record_config

const UPSERT_POST_VOTES_RECORD_COMMAND = `
INSERT INTO post_votes_records
(
  actor,
  post_hash,
  vote_type,
  signed_reputations
)
VALUES 
(
  :actor, 
  :post_hash,
  :vote_type,
  :signed_reputations
);
`

const DELETE_POST_VOTES_RECORDS_BY_ACTOR_AND_POST_HASH_COMMAND = `
DELETE FROM post_votes_records
WHERE actor = $1 and post_hash = $2;
`

const DELETE_POST_VOTES_RECORDS_BY_POST_HASH_COMMAND = `
DELETE FROM post_votes_records
WHERE post_hash = $1;
`

const DELETE_POST_VOTES_RECORDS_BY_ACTOR_COMMAND = `
DELETE FROM post_votes_records
WHERE actor = $1;
`

const QUERY_VOTES_COUNT_BY_VOTE_TYPE_COMMAND = `
SELECT count(*) FROM post_votes_records
WHERE actor = $1 and post_hash = $2 and vote_type = $3;
`

const QUERY_TOTAL_VOTES_COUNT_COMMAND = `
SELECT count(*) FROM post_votes_records
WHERE actor = $1 and post_hash = $2;
`

const QUERY_ACTOR_LIST_BY_POST_HASH_AND_VOTE_TYPE_COMMAND = `
SELECT actor FROM post_votes_records
WHERE post_hash = $1 and vote_type = $2;
`