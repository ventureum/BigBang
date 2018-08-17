package actor_profile_record_config


const UPSERT_ACTOR_PROFILE_RECORD_COMMAND = `
INSERT INTO actor_profile_records
(
  actor,
  actor_type
)
VALUES 
(
  :actor, 
  :actor_type
)
ON CONFLICT (actor) 
DO
 UPDATE
    SET actor_type = :actor_type
    WHERE actor_profile_records.actor = :actor;
`

const DELETE_ACTOR_PROFILE_RECORD_COMMAND = `
DELETE FROM actor_profile_records
WHERE actor = $1;
`

const QUERY_ACTOR_PROFILE_RECORD_COMMAND = `
SELECT * FROM actor_profile_records
WHERE actor = $1;
`
