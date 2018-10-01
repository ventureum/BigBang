package actor_rewards_info_record_config


const UPSERT_ACTOR_REWARDS_INFO_RECORD_COMMAND = `
INSERT INTO actor_rewards_info_records
(
  actor,
  fuel,
  reputation,
  milestone_points
)
VALUES 
(
  :actor,
  :fuel,
  :reputation,
  :milestone_points
)
ON CONFLICT (actor) 
DO
 UPDATE
    SET 
        fuel = :fuel,
        reputation = :reputation,
        milestone_points = :milestone_points
    WHERE actor_rewards_info_records.actor = :actor;
`

const DELETE_ACTOR_REWARDS_INFO_RECORD_COMMAND = `
DELETE FROM actor_rewards_info_records
WHERE actor = $1;
`

const QUERY_ACTOR_REWARDS_INFO_COMMAND = `
SELECT fuel, reputation, milestone_points FROM actor_rewards_info_records
WHERE actor = $1;
`

const ADD_ACTOR_MILESTONE_POINTS_COMMAND = `
UPDATE actor_rewards_info_records
    SET milestone_points = actor_rewards_info_records.milestone_points + $2
    WHERE actor = $1
RETURNING milestone_points;
`

const ADD_ACTOR_REPUTATIONS_COMMAND = `
UPDATE actor_rewards_info_records
    SET reputation = actor_rewards_info_records.reputation + $2
    WHERE actor = $1
RETURNING reputation;
`

const ADD_ACTOR_FUEL_COMMAND = `
UPDATE actor_rewards_info_records
    SET fuel = actor_rewards_info_records.fuel + $2
    WHERE actor = $1
RETURNING fuel;
`

const SUB_ACTOR_REPUTATION_COMMAND = `
UPDATE actor_rewards_info_records x
  SET reputation = LEAST(GREATEST(y.reputation - $2, 0), y.reputation)
  FROM (SELECT actor, reputation FROM actor_rewards_info_records WHERE actor = $1 FOR UPDATE) y 
  WHERE x.actor = y.actor
RETURNING $2 - y.reputation;
`

const SUB_ACTOR_FUEL_COMMAND = `
UPDATE actor_rewards_info_records x
  SET fuel = LEAST(GREATEST(y.fuel - $2, 0), y.fuel)
  FROM (SELECT actor, fuel FROM actor_rewards_info_records WHERE actor = $1 FOR UPDATE) y 
  WHERE x.actor = y.actor
RETURNING $2 - y.fuel;
`

const QUARY_TOTAL_REPUTATION_COMMAND = `
SELECT COALESCE(sum(reputation), 0) FROM actor_rewards_info_records;
`

const VERIFY_ACTOR_EXISTING_COMMAND = `
select exists(select 1 from actor_rewards_info_records where actor =$1);
`
