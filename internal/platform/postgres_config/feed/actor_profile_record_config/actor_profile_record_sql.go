package actor_profile_record_config

const UPSERT_ACTOR_PROFILE_RECORD_COMMAND = `
INSERT INTO actor_profile_records
(
  actor,
  actor_type,
  username,
  photo_url,
  telegram_id,
  phone_number,
  public_key,
  profile_content
)
VALUES 
(
  :actor, 
  :actor_type,
  :username,
  :photo_url,
  :telegram_id,
  :phone_number,
  :public_key,
  :profile_content
)
ON CONFLICT (actor) 
DO
 UPDATE
    SET actor_type = :actor_type,
        username = :username,
        photo_url = :photo_url,
        telegram_id = :telegram_id,
        phone_number = :phone_number,
        public_key = :public_key,
        actor_profile_status = 'ACTIVATED',
        profile_content = :profile_content
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
const VERIFY_ACTOR_NO_ADMIN_TYPE_COMMAND = `
select exists(select 1 from actor_profile_records where actor = $1 and actor_profile_status = 'ACTIVATED' and actor_type != 'ADMIN');
`

const QUERY_ACTOR_TYPE_COMMAND = `
SELECT actor_type FROM actor_profile_records
WHERE actor = $1 and actor_profile_status = 'ACTIVATED';
`

const VERIFY_ACTOR_TYPE_COMMAND = `
select exists(select 1 from actor_profile_records where actor = $1 and actor_type = $2 and actor_profile_status = 'ACTIVATED');
`

const UPDATE_ACTOR_PRIVATE_KEY_COMMAND = `
UPDATE actor_profile_records
  SET private_key = $2
WHERE actor = $1;
`

const QUERY_ACTOR_PRIVATE_KEY_COMMAND = `
SELECT private_key 
FROM actor_profile_records
WHERE actor = $1;
`

const QUERY_ACTOR_UUID_FROM_PRIVATE_KEY_COMMAND = `
SELECT actor
FROM actor_profile_records
WHERE public_key = $1;
`
