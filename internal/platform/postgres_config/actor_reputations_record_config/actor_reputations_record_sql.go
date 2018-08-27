package actor_reputations_record_config


const UPSERT_ACTOR_REPUTATIONS_RECORD_COMMAND = `
INSERT INTO actor_reputations_records
(
  actor,
  reputations
)
VALUES 
(
  :actor, 
  :reputations
)
ON CONFLICT (actor) 
DO
 UPDATE
    SET reputations = :reputations
    WHERE actor_reputations_records.actor = :actor;
`

const DELETE_ACTOR_REPUTATIONS_RECORD_COMMAND = `
DELETE FROM actor_reputations_records
WHERE actor = $1;
`

const QUERY_ACTOR_REPUTATIONS_COMMAND = `
SELECT reputations FROM actor_reputations_records
WHERE actor = $1;
`

const ADD_ACTOR_REPUTATIONS_COMMAND = `
UPDATE actor_reputations_records
    SET reputations = actor_reputations_records.reputations + $2
    WHERE actor = $1
RETURNING reputations;
`

const SUB_ACTOR_REPUTATIONS_COMMAND = `
UPDATE actor_reputations_records
  SET reputations = LEAST(GREATEST(actor_reputations_records.reputations - $2, 0), actor_reputations_records.reputations)
  WHERE actor = $1
RETURNING $2 - actor_reputations_records.reputations;
`

const QUARY_TOTAL_REPUTATIONS_COMMAND = `
SELECT COALESCE(sum(reputations), 0) FROM actor_reputations_records;
`

const VERIFY_ACTOR_EXISTING_COMMAND = `
select exists(select 1 from actor_reputations_records where actor =$1);
`
