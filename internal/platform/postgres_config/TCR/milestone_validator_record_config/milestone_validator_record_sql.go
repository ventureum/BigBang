package milestone_validator_record_config

const UPSERT_MILESTONE_VALIDATOR_RECORD_COMMAND = `
INSERT INTO milestone_validator_records
(
  project_id,
  milestone_id,
  validator
)
VALUES 
(
  :project_id,
  :milestone_id,
  :validator
);
`

const DELETE_ALL_MILESTONE_VALIDATOR_RECORDS_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND = `
DELETE FROM milestone_validator_records
WHERE project_id = $1 and milestone_id = $2;
`

const DELETE_MILESTONE_VALIDATOR_RECORD_BY_IDS_AND_VALIDATOR_COMMAND = `
DELETE FROM milestone_validator_records
WHERE project_id = $1 and milestone_id = $2 and validator = $3;
`

const QUERY_MILESTONE_VALIDATOR_LIST_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND = `
SELECT validator FROM milestone_validator_records
WHERE project_id = $1 and milestone_id = $2;
`

const VERIFY_MILESTONE_VALIDATOR_EXISTING_COMMAND = `
select exists(select 1 from milestone_validator_records where project_id = $1 and milestone_id = $2 and validator = $3);
`
