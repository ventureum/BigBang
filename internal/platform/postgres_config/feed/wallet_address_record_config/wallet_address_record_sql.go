package wallet_address_record_config

const UPSERT_WALLET_ADDRESS_RECORD_COMMAND = `
INSERT INTO wallet_address_records
(
  actor,
  wallet_address
)
VALUES 
(
  :actor, 
  :wallet_address
);
`

const DELETE_ALL_WALLET_ADDRESS_RECORDS_BY_ACTOR_COMMAND = `
DELETE FROM wallet_address_records
WHERE actor = $1;
`

const DELETE_WALLET_ADDRESS_RECORDS_BY_ACTOR_AND_ADDRESS_COMMAND = `
DELETE FROM wallet_address_records
WHERE actor = $1 and wallet_address = $2;
`

const QUERY_WALLET_ADDRESS_LIST_BY_ACTOR_COMMAND = `
SELECT wallet_address FROM wallet_address_records
WHERE actor = $1;
`

const VERIFY_WALLET_ADDRESS_EXISTING_COMMAND = `
select exists(select 1 from wallet_address_records where actor = $1 and wallet_address = $2);
`
