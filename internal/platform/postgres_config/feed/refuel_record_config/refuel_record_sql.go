package refuel_record_config

const UPSERT_REFUEL_RECORD_COMMAND = `
INSERT INTO refuel_records
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
);
`

const DELETE_REFUEL_RECORDS_COMMAND = `
DELETE FROM refuel_records
WHERE actor = $1;
`

const QUERY_REFUEL_RECORDS_COMMAND = `
SELECT * FROM refuel_records
WHERE actor = $1;
`

const QUERY_LATEST_REFUEL_TIME_COMMAND = `
SELECT max(created_at) FROM refuel_records
WHERE actor = $1;
`
