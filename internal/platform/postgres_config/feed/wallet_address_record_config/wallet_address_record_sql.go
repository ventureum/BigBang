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

const DELETE_MULTIPLE_WALLET_ADDRESS_RECORDS_BY_ACTOR_AND_LIST_OF_ADDRESSES_COMMAND = `
DELETE FROM wallet_address_records
WHERE actor = :actor and wallet_address IN (:deleted_addresses);
`

const QUERY_WALLET_ADDRESS_LIST_BY_ACTOR_COMMAND = `
SELECT wallet_address FROM wallet_address_records
WHERE actor = $1;
`
