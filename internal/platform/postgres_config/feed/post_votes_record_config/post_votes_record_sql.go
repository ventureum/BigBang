package post_votes_record_config

const UPSERT_POST_VOTES_RECORD_COMMAND = `
INSERT INTO post_votes_records
(
  actor,
  post_hash,
  post_type,
  vote_type,
  delta_fuel,
  delta_reputation,
  delta_milestone_points,
  signed_reputation
)
VALUES 
(
  :actor, 
  :post_hash,
  :post_type,
  :vote_type,
  :delta_fuel,
  :delta_reputation,
  :delta_milestone_points,
  :signed_reputation
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

const QUERY_RECENT_POST_VOTES_RECORDS_BY_ACTOR_COMMAND = `
SELECT * FROM post_votes_records
WHERE actor = $1 ORDER BY created_at DESC LIMIT $2;
`

const ADD_POST_VOTE_DELTA_REWARDS_INFO_COMMAND = `
UPDATE  post_votes_records
  SET delta_fuel =  post_votes_records.delta_fuel + $4,
  delta_reputation =  post_votes_records.delta_reputation + $5,
  delta_milestone_points = post_votes_records.delta_milestone_points + $6
  WHERE actor = $1 and post_hash = $2 and vote_type = $3;
`

const QUERY_TOTAL_REPUTATION_BY_POST_HASH_AND_VOTE_TYPE_WITH_TIME_CUTOFF_COMMAND = `
SELECT sum(abs(signed_reputation)) FROM post_votes_records
WHERE post_hash = $1 and vote_type = $2 and created_at < $4;
`

const QUERY_TOTAL_REPUTATION_BY_POST_HASH_AND_VOTE_TYPE_COMMAND = `
SELECT sum(abs(signed_reputation)) FROM post_votes_records
WHERE post_hash = $1 and vote_type = $2;
`
