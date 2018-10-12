package actor_rating_vote_account_config

const UPSERT_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND = `
INSERT INTO actor_rating_vote_accounts 
(
  actor,
  project_id,
  available_rating_votes,
  received_rating_votes
)
VALUES 
(
  :actor,
  :project_id,
  :available_rating_votes,
  :received_rating_votes
)
ON CONFLICT (actor, project_id) 
DO
 UPDATE
    SET
        available_rating_votes = :available_rating_votes,
        received_rating_votes = :received_rating_votes
    WHERE actor_rating_vote_accounts.actor = :actor;
`

const DELETE_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND = `
DELETE FROM actor_rating_vote_accounts
WHERE actor = $1;
`

const QUERY_ACTOR_RATING_VOTE_ACCOUNT_RECORD_COMMAND = `
SELECT * FROM actor_rating_vote_accounts
WHERE actor = $1;
`
