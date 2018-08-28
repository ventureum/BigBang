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
    SET actor_type = :actor_type,
        actor_profile_status = 'ACTIVATED'
    WHERE actor_profile_records.actor = :actor
RETURNING (xmax = 0) AS inserted;
`

const DELETE_ACTOR_PROFILE_RECORD_COMMAND = `
DELETE FROM actor_profile_records
WHERE actor = $1;
`

const DEACTIVATE_ACTOR_PROFILE_RECORD_COMMAND = `
UPDATE actor_profile_records
  SET actor_profile_status = 'DEACTIVATED'
WHERE actor = $1;
`

const QUERY_ACTOR_PROFILE_RECORD_COMMAND = `
SELECT * FROM actor_profile_records
WHERE actor = $1;
`

const VERIFY_ACTOR_EXISTING_COMMAND = `
select exists(select 1 from actor_profile_records where actor = $1 and actor_profile_status = 'ACTIVATED');
`
