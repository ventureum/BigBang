package proxy_config

const UPSERT_PROXY_COMMAND = `
INSERT INTO proxies 
(
 uuid
)
VALUES 
(
 :uuid
)
ON CONFLICT(uuid) 
DO NOTHING;
`

const DELETE_PROXY_COMMAND = `
DELETE FROM proxies
WHERE uuid = $1;
`

const QUERY_PROXY_COMMAND = `
SELECT * FROM proxies WHERE uuid = $1;
`

const QUERY_LIST_OF_PROXY_UUIDS_COMMAND = `
SELECT uuid FROM proxies;
`

const VERIFY_PROXY_EXISTING_COMMAND = `
select exists(select 1 from proxies where uuid =$1);
`
