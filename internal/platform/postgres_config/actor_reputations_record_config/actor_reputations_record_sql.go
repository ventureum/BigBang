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

const QUERY_ACTOR_REPUTATIONS_FOR_UPDATE_COMMAND = `
SELECT reputations FROM actor_reputations_records
WHERE actor = $1 FOR UPDATE;
`

const ADD_ACTOR_REPUTATIONS_COMMAND = `
INSERT INTO actor_reputations_records
(
  actor,
  reputations
)
VALUES 
(
  $1, 
  $2
)
ON CONFLICT (actor) 
DO
 UPDATE
    SET reputations = actor_reputations_records.reputations + $2
    WHERE actor_reputations_records.actor = $1;
`

const SUB_ACTOR_REPUTATIONS_COMMAND = `
UPDATE actor_reputations_records
  SET reputations = actor_reputations_records.reputations - $2
  WHERE actor = $1;
`

const QUARY_TOTAL_REPUTATIONS_COMMAND = `
SELECT sum(reputations) FROM actor_reputations_records;
`
