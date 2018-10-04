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

const PAGINATION_QUERY_LIST_OF_PROXY_COMMAND = `
SELECT * FROM proxies
WHERE id <= $1
ORDER BY id DESC
LIMIT $2;
`

const QUERY_LIST_OF_PROXY_COMMAND = `
SELECT * FROM proxies
ORDER BY id DESC
LIMIT $1;
`

const VERIFY_PROXY_EXISTING_COMMAND = `
select exists(select 1 from proxies where uuid =$1);
`
