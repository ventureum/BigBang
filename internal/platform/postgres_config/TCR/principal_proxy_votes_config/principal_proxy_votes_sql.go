package principal_proxy_votes_config


const INSERT_PRINCIPAL_PROXY_VOTES_COMMAND = `
INSERT INTO principal_proxy_votes 
( 
  id,
  actor,
  project_id,
  proxy,
  block_timestamp,
  votes
)
VALUES 
(
  :id,
  :actor,
  :project_id,
  :proxy,
  :block_timestamp,
  :votes
);
`

const UPDATE_PRINCIPAL_PROXY_VOTES_COMMAND = `
UPDATE principal_proxy_votes
    SET
          block_timestamp = :block_timestamp,
          votes = :votes
    WHERE 
          principal_proxy_votes.actor = :actor and
          principal_proxy_votes.project_id = :project_id and
          principal_proxy_votes.proxy = :proxy;
`

const DELETE_PRINCIPAL_PROXY_VOTES_BY_IDS_COMMAND = `
DELETE FROM principal_proxy_votes
WHERE actor = $1 and project_id = $2 and proxy = $3;
`

const QUERY_PRINCIPAL_PROXY_VOTES_BY_IDS_COMMAND = `
SELECT * FROM principal_proxy_votes
WHERE actor = $1 and project_id = $2 and proxy = $3;
`

const VERIFY_PRINCIPAL_PROXY_VOTES_EXISTING_COMMAND = `
select exists(select 1 from principal_proxy_votes where actor = $1 and project_id = $2 and proxy = $3);
`

const PAGINATION_QUERY_PRINCIPAL_PROXY_VOTES_LIST_COMMAND = `
SELECT * FROM principal_proxy_votes
WHERE id <= $3 and actor = $1 and project_id = $2
ORDER BY id DESC
LIMIT $4;
`

const QUERY_PRINCIPAL_PROXY_VOTES_LIST_COMMAND = `
SELECT * FROM principal_proxy_votes
WHERE actor = $1 and project_id = $2
ORDER BY id DESC
LIMIT $3;
`