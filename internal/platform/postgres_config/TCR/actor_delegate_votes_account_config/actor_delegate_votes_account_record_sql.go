package actor_delegate_votes_account_config

const UPSERT_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND = `
INSERT INTO actor_delegate_votes_accounts 
(
  actor,
  project_id,
  available_delegate_votes,
  received_delegate_votes
)
VALUES 
(
  :actor,
  :project_id,
  :available_delegate_votes,
  :received_delegate_votes
)
ON CONFLICT (actor, project_id) 
DO
 UPDATE
    SET
        available_delegate_votes = :available_delegate_votes,
        received_delegate_votes = :received_delegate_votes
    WHERE actor_delegate_votes_accounts.actor = :actor;
`

const DELETE_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND = `
DELETE FROM actor_delegate_votes_accounts
WHERE actor = $1 and project_id = $2;
`

const QUERY_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND = `
SELECT * FROM actor_delegate_votes_accounts
WHERE actor = $1 and project_id = $2;
`

const UPDATE_AVAILABLE_DELEGATE_VOTES_COMMAND = `
UPDATE  actor_delegate_votes_accounts
SET available_delegate_votes = $3
WHERE actor = $1 and project_id = $2;
`

const UPDATE_RECEIVED_DELEGATE_VOTES_COMMAND = `
UPDATE  actor_delegate_votes_accounts
SET received_delegate_votes = received_delegate_votes + $3
WHERE actor = $1 and project_id = $2;
`

const VERIFY_DELEGATE_VOTES_ACCOUNT_EXISTING_COMMAND = `
select exists(select 1 from actor_delegate_votes_accounts where actor = $1 and project_id = $2);
`
