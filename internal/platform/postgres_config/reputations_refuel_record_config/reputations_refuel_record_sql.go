package reputations_refuel_record_config


const UPSERT_REPUTATION_REFUEL_RECORD_COMMAND = `
INSERT INTO reputations_refuel_records
(
  actor,
  reputations
)
VALUES 
(
  :actor, 
  :reputations
);
`

const DELETE_REPUTATION_REFUEL_RECORDS_COMMAND = `
DELETE FROM reputations_refuel_records
WHERE actor = $1;
`

const QUERY_REPUTATION_REFUEL_RECORDS_COMMAND = `
SELECT * FROM reputations_refuel_records
WHERE actor = $1;
`
